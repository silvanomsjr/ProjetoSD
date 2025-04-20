# Projeto KVS - Sistemas Distribuidos

## Uso

### Pre-requisitos

- Go 1.23.4 ou superior
- Mosquitto MQTT broker rodando no localhost:1883

### Build do Projeto

```bash
# Gera o código gRPC do protobuf
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/kvs.proto

# Faz o build do projeto todo
go build
```


## Rodando

Antes de tudo, devemos possuir o "mosquitto" já em execução na sua porta padrão (1883).

Então antes de tudo, rodar:

```bash
mosquitto
```

Logo após, já podemos observar as flags existentes no projeto buildando com a flag "--help":


```bash
./ProjetoKVS --help

Usage of ./ProjetoKVS:
  -mqtt_broker string
    	Endereço do broker MQTT (default "tcp://localhost:1883")
  -mqtt_client_id string
    	ID do cliente MQTT (padrão: porta do servidor)
  -mqtt_topic string
    	Tópico MQTT para atualizações (default "kvs_updates")
  -port int
    	A porta do servidor (default 50051)
```


Para rodar de forma que passe em todos os testes, devemos abrir 3, terminais (ou utilizar Tmux).

E em cada um deles, rodar:

```bash
./ProjetoKVS --port="9000"

./ProjetoKVS --port="9001"

./ProjetoKVS --port="9002"
```


Após isso, devemos utilizar o código de "kvs-client-2024-2" para que rodemos o cliente em Rust, de tal maneira, o projeto já estara 100% funcional e implementando todas as tecnologias necessárias.
