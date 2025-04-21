#!/bin/bash

set -e

# Mostra etapas de execução
echo "=== Instalando dependências e compilando servidor KVS ==="

# Verifica se o Go está instalado
if ! command -v go &> /dev/null; then
    echo "Erro: Go não está instalado. Por favor, instale Go 1.17 ou posterior."
    exit 1
fi

# Verifica se o protoc está instalado
if ! command -v protoc &> /dev/null; then
    echo "Erro: protoc não está instalado. Por favor, instale o compilador Protocol Buffers."
    echo "Visite https://github.com/protocolbuffers/protobuf/releases para instruções de instalação."
    exit 1
fi

# Instala pacotes Go necessários
echo "Instalando dependências Go..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Atualiza módulos Go
echo "Atualizando módulos Go..."
go mod tidy

# Gera código Protocol Buffer e gRPC
echo "Gerando código Protocol Buffer..."
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/kvs.proto

# Compila o servidor
echo "Compilando binário do servidor..."
go build -o ProjetoKVS main.go

# Verifica se a compilação foi bem-sucedida
if [ -f "ProjetoKVS" ]; then
    echo "=== Compilação concluída com sucesso! ==="
    echo "Você pode agora executar o servidor usando: ./server.sh <porta>"
else
    echo "=== Compilação falhou! ==="
fi
