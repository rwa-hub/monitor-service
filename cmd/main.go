package main

import (
	"log"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/internal/app/service"
	"net/http"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Ocorreu um panic: %v", r)
		}
	}()

	log.Println("Iniciando o serviço de monitoramento de contratos inteligentes...")

	rpcURL := "http://127.0.0.1:8545"
	contractAddress := "0x9E545E3C0baAB3E08CdfD552C960A1050f373042" //compliance

	// Cria uma nova instância do cliente RPC
	rpcClient, err := rpc.NewRPCClient(rpcURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao nó Ethereum: %v", err)
	}

	wsServer := websocket.NewWebSocketServer()

	// WebSocket
	http.HandleFunc("/ws", wsServer.HandleConnections)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// start
	service.EventService(rpcClient, wsServer, contractAddress)

	select {}
}
