package main

import (
	"log"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/internal/app/service"
)

func main() {
	rpcURL := "ws://host.docker.internal:8546"
	contractAddress := "0xFa5B6432308d45B54A1CE1373513Fab77166436f"

	wsServer := websocket.NewWebSocketServer()
	go wsServer.HandleMessages()

	rpcClient, err := rpc.NewRPCClient(rpcURL)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao nó Ethereum: %v", err)
	}
	defer func() {
		client, err := rpcClient.Client()
		if err == nil {
			client.Close()
		}
	}()

	service.EventService(rpcClient, wsServer, contractAddress)

	select {}
}
