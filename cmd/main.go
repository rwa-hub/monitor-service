package main

import (
	"log"
	"net/http"

	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	service "monitor-service/internal/service"
)

func main() {
	rpcURL := "ws://host.docker.internal:8546"
	contractAddresses := map[string]string{
		"financial_compliance":   "0x486B9f3Fb089f9166147149562e0A6952442Fa42",
		"modular_compliance":     "0x9F035Be9853eb3Fc10155361fda197C159eFDD09",
		"ident_registry_storage": "0xf935973e9f884c66e6E3ce681344064e565e0250",
		"registry_md":            "0xC8c03647d39a96f02f6Ce8999bc22493C290e734",
	}

	// Inicializa o servidor WebSocket
	wsServer := websocket.NewWebSocketServer()
	http.HandleFunc("/ws", wsServer.HandleConnections)

	// Inicializa a conexão RPC
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

	// Inicia os serviços de eventos
	go service.EventService(rpcClient, wsServer, contractAddresses)

	// Define a porta do WebSocket
	port := ":8080"
	log.Printf("✅ Servidor WebSocket iniciado na porta %s", port)

	// Inicia o servidor HTTP para WebSocket
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("❌ Erro ao iniciar o servidor HTTP: %v", err)
	}
}
