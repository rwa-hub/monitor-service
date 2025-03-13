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
		"financial_compliance":   "0xF235Bb30Ad375F279248AaFC89F4A92899A900De",
		"modular_compliance":     "0x850F95B0f32E9dB5AA484d160CB58f8A52103dc2",
		"ident_registry_storage": "0xd9bB4402537D044709BC80b666F522A6F0ce6435",
		"registry_md":            "0xC8c03647d39a96f02f6Ce8999bc22493C290e734",
		"identity_registry":      "0x9b5b3cCe1f7f8359E026f9573f258782Be577f29",
		"token_rwa":              "0x2078FaF714Fb3727a66bc10F7A9690b5A16CD0bb",
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
