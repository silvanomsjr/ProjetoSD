#!/bin/bash

# Verifica se o parâmetro de porta foi fornecido
if [ $# -ne 1 ]; then
    echo "Uso: $0 <porta>"
    echo "Exemplo: $0 50051"
    exit 1
fi

PORTA=$1

# Valida que a porta é um número
if ! [[ "$PORTA" =~ ^[0-9]+$ ]]; then
    echo "Erro: A porta deve ser um número"
    exit 1
fi

# Verifica se o binário existe
if [ ! -f "ProjetoKVS" ]; then
    echo "Erro: Binário do servidor não encontrado. Execute ./compile.sh primeiro."
    exit 1
fi

# Define o ID do cliente MQTT baseado na porta
MQTT_CLIENT_ID="kvs-server-$PORTA"

echo "Iniciando servidor KVS na porta $PORTA com ID de cliente MQTT $MQTT_CLIENT_ID"
echo "Pressione Ctrl+C para parar o servidor"

# Executa o servidor com a porta especificada
./ProjetoKVS --port="$PORTA" --mqtt_client_id="$MQTT_CLIENT_ID"
