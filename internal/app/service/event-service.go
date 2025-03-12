package service

import (
	"encoding/json"
	"log"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
)

// EventService gerencia a lógica de eventos
func EventService(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, contractAddress string, eventSignature string) {
	go wsServer.HandleMessages()

	rpcClient.ListenEvents(contractAddress, eventSignature)

	for {
		select {
		case event := <-rpcClient.EventChannel():
			log.Println("Enviando evento para WebSocket:", event)
			eventBytes, err := json.Marshal(event)
			if err != nil {
				log.Println("Erro ao converter evento para JSON:", err)
				continue
			}
			wsServer.Broadcast(eventBytes)
		}
	}
}
