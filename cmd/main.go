package main

import (
	"log"
	"net/http"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	service "monitor-service/internal/service"
)

func main() {
	rpcURL := "ws://host.docker.internal:8546"
	mongoURI := "mongodb://admin:password@host.docker.internal:27017/"
	mongoDBName := "monitor-service"

	contractAddresses := map[string]string{
		"financial_compliance":   "0x486B9f3Fb089f9166147149562e0A6952442Fa42",
		"modular_compliance":     "0x9F035Be9853eb3Fc10155361fda197C159eFDD09",
		"ident_registry_storage": "0xf935973e9f884c66e6E3ce681344064e565e0250",
		"registry_md":            "0xC8c03647d39a96f02f6Ce8999bc22493C290e734",
		"identity_registry":      "0x6EC5189503e0F03704B737d1977230c5b800A7F5",
	}

	wsServer := websocket.NewWebSocketServer()
	http.HandleFunc("/ws", wsServer.HandleConnections)

	rpcClient, err := rpc.NewRPCClient(rpcURL)
	if err != nil {
		log.Fatalf("❌ Error the connection to the RPC: %v", err)
	}

	client, err := rpcClient.Client()
	if err == nil {
		defer client.Close()
	}

	db, err := database.NewMongoDB(mongoURI, mongoDBName)
	if err != nil {
		log.Fatalf("❌ Error the connection to the MongoDB: %v", err)
	}

	queueService, err := queue.NewRabbitMQ("monitor-events")
	if err != nil {
		log.Fatalf("❌ Error the connection to the RabbitMQ: %v", err)
	}
	defer queueService.Close()

	go service.EventService(rpcClient, wsServer, queueService, db, contractAddresses)

	port := ":8080"
	log.Printf("✅ Server WebSocket started on port %s", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("❌ Error to start the HTTP server: %v", err)
	}
}
