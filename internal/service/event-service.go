package listeners

import (
	"log"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"

	"monitor-service/internal/service/listeners/financial_compliance"
	"monitor-service/internal/service/listeners/ident_registry_storage"
	"monitor-service/internal/service/listeners/identity_registry"
	"monitor-service/internal/service/listeners/modular_compliance"
	"monitor-service/internal/service/listeners/registry_md"
	"monitor-service/internal/service/listeners/token_rwa"
)

func EventService(
	rpcClient *rpc.RPCClient,
	wsServer *websocket.WebSocketServer,
	queueService *queue.RabbitMQ,
	db *database.MongoDB,
	contractAddresses map[string]string,
) {
	log.Println("🚀 Starting event service...")

	go func() {
		defer recoverPanic("WebSocket Server")
		wsServer.HandleMessages()
	}()

	contractMappings := map[string]func(*rpc.RPCClient, *websocket.WebSocketServer, *queue.RabbitMQ, *database.MongoDB, string){
		"financial_compliance":   financial_compliance.StartListener,
		"ident_registry_storage": ident_registry_storage.StartListener,
		"modular_compliance":     modular_compliance.StartListener,
		"registry_md":            registry_md.StartListener,
		"identity_registry":      identity_registry.StartListener,
		"token_rwa":              token_rwa.StartListener,
	}

	for key, startFunc := range contractMappings {
		if address, exists := contractAddresses[key]; exists && address != "" {
			go func(k, addr string, fn func(*rpc.RPCClient, *websocket.WebSocketServer, *queue.RabbitMQ, *database.MongoDB, string)) {
				defer recoverPanic("Listener: " + k)
				log.Printf("✅ Starting listener for %s in contract %s\n", k, addr)
				fn(rpcClient, wsServer, queueService, db, addr)
			}(key, address, startFunc)
		} else {
			log.Printf("⚠️  Contract address for %s not found. Listener will not be started.", key)
		}
	}

	log.Println("✅ All event listeners have been started!")
}

func recoverPanic(serviceName string) {
	if r := recover(); r != nil {
		log.Printf("❌ Panic recovered in %s: %v", serviceName, r)
	}
}
