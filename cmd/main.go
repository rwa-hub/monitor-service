package main

import (
	"log"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/internal/app/service"
	"net/http"
)

func main() {
	log.Println("Iniciando o serviço de monitoramento de contratos inteligentes...")

	rpcURL := "http://127.0.0.1:8545"
	contractAddress := "0x0000000000000000000000000000000000000000"
	eventSignature := "SeuEventoAssinatura"

	rpcClient, err := rpc.NewRPCClient(rpcURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao nó Ethereum: %v", err)
	}

	wsServer := websocket.NewWebSocketServer()
	http.HandleFunc("/ws", wsServer.HandleConnections)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	service.EventService(rpcClient, wsServer, contractAddress, eventSignature)
}
