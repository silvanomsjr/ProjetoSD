package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "ProjetoKVS/proto"
)

var (
	port         = flag.Int("port", 50051, "A porta do servidor")
	mqttBroker   = flag.String("mqtt_broker", "tcp://localhost:1883", "Endereço do broker MQTT")
	mqttTopic    = flag.String("mqtt_topic", "kvs_updates", "Tópico MQTT para atualizações")
	mqttClientID = flag.String("mqtt_client_id", "", "ID do cliente MQTT (padrão: porta do servidor)")
)

type KVSServer struct {
	pb.UnimplementedKVSServer
	mu          sync.RWMutex
	store       map[string]map[int32]string
	maxVersions map[string]int32
	mqttClient  mqtt.Client
}

type MQTTOperation string

const (
	InsertOperation MQTTOperation = "INSERT"
	RemoveOperation MQTTOperation = "REMOVE"
)

type MQTTMessage struct {
	Operation MQTTOperation `json:"operation"`
	Key       string        `json:"key"`
	Value     string        `json:"value,omitempty"`
	Version   int32         `json:"version"`
	ServerID  string        `json:"server_id"`
}

func NewKVSServer(mqttClientID string) *KVSServer {
	server := &KVSServer{
		store:       make(map[string]map[int32]string),
		maxVersions: make(map[string]int32),
	}

	opts := mqtt.NewClientOptions().
		AddBroker(*mqttBroker).
		SetClientID(mqttClientID).
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetConnectionLostHandler(func(client mqtt.Client, err error) {
			log.Printf("MQTT Connection lost: %v", err)
		}).
		SetReconnectingHandler(func(client mqtt.Client, opts *mqtt.ClientOptions) {
			log.Println("MQTT Reconnecting...")
		})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	token := client.Subscribe(*mqttTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		server.handleMQTTMessage(msg.Payload())
	})
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to MQTT topic: %v", token.Error())
	}

	server.mqttClient = client
	return server
}

func (s *KVSServer) publishUpdate(op MQTTOperation, key string, value string, version int32) {
	msg := MQTTMessage{
		Operation: op,
		Key:       key,
		Value:     value,
		Version:   version,
		ServerID:  *mqttClientID,
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal MQTT message: %v", err)
		return
	}

	token := s.mqttClient.Publish(*mqttTopic, 0, false, payload)
	if token.Wait() && token.Error() != nil {
		log.Printf("Failed to publish MQTT message: %v", token.Error())
	}
}

func (s *KVSServer) handleMQTTMessage(payload []byte) {
	var msg MQTTMessage
	if err := json.Unmarshal(payload, &msg); err != nil {
		log.Printf("Failed to unmarshal MQTT message: %v", err)
		return
	}

	if msg.ServerID == *mqttClientID {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	switch msg.Operation {
	case InsertOperation:
		if _, exists := s.store[msg.Key]; !exists {
			s.store[msg.Key] = make(map[int32]string)
		}

		s.store[msg.Key][msg.Version] = msg.Value
		if msg.Version > s.maxVersions[msg.Key] {
			s.maxVersions[msg.Key] = msg.Version
		}

	case RemoveOperation:
		if _, exists := s.store[msg.Key]; exists {
			if msg.Version == 0 {
				delete(s.store, msg.Key)
			} else {
				delete(s.store[msg.Key], msg.Version)

				if len(s.store[msg.Key]) == 0 {
					delete(s.store, msg.Key)
				}
			}
		}
	}
}

func validateKeyValue(key, value string) error {
	if len(key) < 3 {
		return status.Errorf(codes.InvalidArgument, "key must have at least 3 characters")
	}
	if len(value) < 3 {
		return status.Errorf(codes.InvalidArgument, "value must have at least 3 characters")
	}
	return nil
}

func (s *KVSServer) Insere(ctx context.Context, in *pb.ChaveValor) (*pb.Versao, error) {
	if err := validateKeyValue(in.Chave, in.Valor); err != nil {
		return &pb.Versao{Versao: -1}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	key := in.Chave
	value := in.Valor

	if _, exists := s.store[key]; !exists {
		s.store[key] = make(map[int32]string)
	}

	nextVersion := s.maxVersions[key] + 1
	if nextVersion == 0 {
		nextVersion = 1
	}

	s.store[key][nextVersion] = value
	s.maxVersions[key] = nextVersion

	s.publishUpdate(InsertOperation, key, value, nextVersion)

	return &pb.Versao{Versao: nextVersion}, nil
}

func (s *KVSServer) Consulta(ctx context.Context, in *pb.ChaveVersao) (*pb.Tupla, error) {
	key := in.Chave

	if len(key) < 3 {
		return &pb.Tupla{Chave: "", Valor: "", Versao: 0}, status.Errorf(codes.InvalidArgument, "A key deve ter mais de 3 caracteres")
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	versions, exists := s.store[key]
	if !exists || len(versions) == 0 {
		return &pb.Tupla{Chave: "", Valor: "", Versao: 0}, nil
	}

	var targetVersion int32
	if in.Versao != nil && *in.Versao > 0 {
		targetVersion = *in.Versao
	} else {
		targetVersion = s.maxVersions[key]
	}

	var highestVersion int32
	for ver := range versions {
		if ver <= targetVersion && ver > highestVersion {
			highestVersion = ver
		}
	}

	if highestVersion == 0 {
		return &pb.Tupla{Chave: "", Valor: "", Versao: 0}, nil
	}

	return &pb.Tupla{
		Chave:  key,
		Valor:  versions[highestVersion],
		Versao: highestVersion,
	}, nil
}

func (s *KVSServer) Remove(ctx context.Context, in *pb.ChaveVersao) (*pb.Versao, error) {
	key := in.Chave

	if len(key) < 3 {
		return &pb.Versao{Versao: -1}, status.Errorf(codes.InvalidArgument, "key must have at least 3 characters")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	versions, exists := s.store[key]
	if !exists || len(versions) == 0 {
		return &pb.Versao{Versao: -1}, nil
	}

	var removedVersion int32 = -1

	if in.Versao != nil {
		ver := *in.Versao
		if _, exists := versions[ver]; exists {
			delete(versions, ver)
			removedVersion = ver

			if len(versions) == 0 {
				delete(s.store, key)
			}
		} else {
			return &pb.Versao{Versao: -1}, nil
		}
	} else {
		removedVersion = 0
		delete(s.store, key)
	}

	s.publishUpdate(RemoveOperation, key, "", removedVersion)

	return &pb.Versao{Versao: removedVersion}, nil
}

func (s *KVSServer) InsereVarias(stream pb.KVS_InsereVariasServer) error {
	var results []*pb.Versao

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		result, err := s.Insere(stream.Context(), in)
		if err != nil {
			results = append(results, &pb.Versao{Versao: -1})
		} else {
			results = append(results, result)
		}
	}

	for _, result := range results {
		if err := stream.Send(result); err != nil {
			return err
		}
	}

	return nil
}

func (s *KVSServer) ConsultaVarias(stream pb.KVS_ConsultaVariasServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		result, err := s.Consulta(stream.Context(), in)
		if err != nil {
			if sendErr := stream.Send(&pb.Tupla{Chave: "", Valor: "", Versao: 0}); sendErr != nil {
				return sendErr
			}
		} else {
			if sendErr := stream.Send(result); sendErr != nil {
				return sendErr
			}
		}
	}
}

func (s *KVSServer) RemoveVarias(stream pb.KVS_RemoveVariasServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		result, err := s.Remove(stream.Context(), in)
		if err != nil {
			if sendErr := stream.Send(&pb.Versao{Versao: -1}); sendErr != nil {
				return sendErr
			}
		} else {
			if sendErr := stream.Send(result); sendErr != nil {
				return sendErr
			}
		}
	}
}

func (s *KVSServer) Snapshot(in *pb.Versao, stream pb.KVS_SnapshotServer) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	targetVersion := in.Versao

	useLatest := targetVersion <= 0

	for key, versions := range s.store {
		if len(versions) == 0 {
			continue
		}

		var highestVersion int32
		var value string

		if useLatest {
			highestVersion = s.maxVersions[key]
			value = versions[highestVersion]
		} else {

			for ver, val := range versions {
				if ver <= targetVersion && ver > highestVersion {
					highestVersion = ver
					value = val
				}
			}
		}

		if highestVersion > 0 {
			if err := stream.Send(&pb.Tupla{
				Chave:  key,
				Valor:  value,
				Versao: highestVersion,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()

	if *mqttClientID == "" {
		*mqttClientID = fmt.Sprintf("kvs-server-%d", *port)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kvsServer := NewKVSServer(*mqttClientID)
	pb.RegisterKVSServer(grpcServer, kvsServer)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down server...")

		if kvsServer.mqttClient != nil && kvsServer.mqttClient.IsConnected() {
			kvsServer.mqttClient.Disconnect(250)
		}

		grpcServer.GracefulStop()
		log.Println("Server stopped")
		os.Exit(0)
	}()

	log.Printf("Server started on port %d", *port)
	log.Printf("Connected to MQTT broker at %s", *mqttBroker)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
