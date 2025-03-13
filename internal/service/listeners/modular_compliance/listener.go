package modular_compliance

import (
	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/utils"

	modularcompliance "monitor-service/internal/modules/ModularCompliance"

	"time"

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
			utils.ProcessEvent("ModuleAdded", event, wsServer, queueService, db, "modular_compliance")
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
			utils.ProcessEvent("ModuleRemoved", event, wsServer, queueService, db, "modular_compliance")
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
			utils.ProcessEvent("OwnershipTransferred", event, wsServer, queueService, db, "modular_compliance")
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
			utils.ProcessEvent("TokenBound", event, wsServer, queueService, db, "modular_compliance")
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
			utils.ProcessEvent("TokenUnbound", event, wsServer, queueService, db, "modular_compliance")
		}

		sub.Unsubscribe()
	}
}
