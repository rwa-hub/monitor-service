package modular_compliance

import (
	"encoding/json"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	modularcompliance "monitor-service/internal/modules/ModularCompliance"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error creating ModularCompliance contract instance")
		return
	}

	go listenForModuleAdded(contract, wsServer, queueService, db)
	go listenForModuleRemoved(contract, wsServer, queueService, db)
	go listenForOwnershipTransferred(contract, wsServer, queueService, db)
	go listenForTokenBound(contract, wsServer, queueService, db)
	go listenForTokenUnbound(contract, wsServer, queueService, db)
}

func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*modularcompliance.Modularcompliance, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return modularcompliance.NewModularcompliance(common.HexToAddress(contractAddress), client)
}

func processEvent(eventType string, event interface{}, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Error converting event to JSON")
		return
	}

	wsServer.Broadcast(eventBytes)

	err = queueService.Publish(eventBytes)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Error sending event to RabbitMQ")
	}

	err = db.InsertEvent(eventType, event)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Error storing event in MongoDB")
	}

	logger.Log.Info().Str("eventType", eventType).Msg("📢 Event processed successfully")
}

func listenForModuleAdded(contract *modularcompliance.Modularcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *modularcompliance.ModularcomplianceModuleAdded)

		sub, err := contract.WatchModuleAdded(opts, eventCh, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Erro ao escutar eventos ModuleAdded. Tentando novamente em 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("ModuleAdded", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForModuleRemoved(contract *modularcompliance.Modularcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *modularcompliance.ModularcomplianceModuleRemoved)

		sub, err := contract.WatchModuleRemoved(opts, eventCh, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening for ModuleRemoved events. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("ModuleRemoved", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForOwnershipTransferred(contract *modularcompliance.Modularcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *modularcompliance.ModularcomplianceOwnershipTransferred)

		sub, err := contract.WatchOwnershipTransferred(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening for OwnershipTransferred events. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("OwnershipTransferred", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForTokenBound(contract *modularcompliance.Modularcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *modularcompliance.ModularcomplianceTokenBound)

		sub, err := contract.WatchTokenBound(opts, eventCh)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening for TokenBound events. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("TokenBound", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForTokenUnbound(contract *modularcompliance.Modularcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *modularcompliance.ModularcomplianceTokenUnbound)

		sub, err := contract.WatchTokenUnbound(opts, eventCh)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening for TokenUnbound events. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("TokenUnbound", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}
