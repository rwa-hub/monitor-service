package registry_md

import (
	"encoding/json"
	"math/big"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	registrymdcompliance "monitor-service/internal/modules/compliance/RegistryMD"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error to initialize RegistryMDCompliance contract")
		return
	}

	go listenForAverbacaoAdded(contract, wsServer, queueService, db)
	go listenForComplianceBound(contract, wsServer, queueService, db)
	go listenForComplianceUnbound(contract, wsServer, queueService, db)
	go listenForOwnershipTransferred(contract, wsServer, queueService, db)
	go listenForPropertyRegistered(contract, wsServer, queueService, db)
	go listenForPropertyTransferred(contract, wsServer, queueService, db)
	go listenForPropertyUpdated(contract, wsServer, queueService, db)
}

func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*registrymdcompliance.Registrymdcompliance, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return registrymdcompliance.NewRegistrymdcompliance(common.HexToAddress(contractAddress), client)
}

func listenForEvent[T any](
	eventName string,
	watchFunc func(*bind.WatchOpts, chan<- *T) (event.Subscription, error),
	wsServer *websocket.WebSocketServer,
	queueService *queue.RabbitMQ,
	db *database.MongoDB,
) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *T)

		sub, err := watchFunc(opts, eventCh)
		if err != nil {
			logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error to listen event. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent(eventName, event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func processEvent(eventName string, event interface{}, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error to convert event to JSON")
		return
	}

	logger.Log.Info().Str("event", eventName).RawJSON("data", eventBytes).Msg("📢 Event captured")
	wsServer.Broadcast(eventBytes)

	if err := queueService.Publish(eventBytes); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error to publish event in queue")
	}

	if err := db.InsertEvent("events", event); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error to save event in MongoDB")
	}
}

func listenForAverbacaoAdded(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("AverbacaoAdded", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcomplianceAverbacaoAdded) (event.Subscription, error) {
		return contract.WatchAverbacaoAdded(opts, sink, []*big.Int{})
	}, wsServer, queueService, db)
}

func listenForComplianceBound(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("ComplianceBound", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcomplianceComplianceBound) (event.Subscription, error) {
		return contract.WatchComplianceBound(opts, sink, []common.Address{})
	}, wsServer, queueService, db)
}

func listenForComplianceUnbound(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("ComplianceUnbound", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcomplianceComplianceUnbound) (event.Subscription, error) {
		return contract.WatchComplianceUnbound(opts, sink, []common.Address{})
	}, wsServer, queueService, db)
}

func listenForOwnershipTransferred(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("OwnershipTransferred", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcomplianceOwnershipTransferred) (event.Subscription, error) {
		return contract.WatchOwnershipTransferred(opts, sink, []common.Address{}, []common.Address{})
	}, wsServer, queueService, db)
}

func listenForPropertyRegistered(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("PropertyRegistered", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcompliancePropertyRegistered) (event.Subscription, error) {
		return contract.WatchPropertyRegistered(opts, sink, []*big.Int{}, []common.Address{})
	}, wsServer, queueService, db)
}

func listenForPropertyTransferred(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("PropertyTransferred", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcompliancePropertyTransferred) (event.Subscription, error) {
		return contract.WatchPropertyTransferred(opts, sink, []*big.Int{}, []common.Address{}, []common.Address{})
	}, wsServer, queueService, db)
}

func listenForPropertyUpdated(contract *registrymdcompliance.Registrymdcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	listenForEvent("PropertyUpdated", func(opts *bind.WatchOpts, sink chan<- *registrymdcompliance.RegistrymdcompliancePropertyUpdated) (event.Subscription, error) {
		return contract.WatchPropertyUpdated(opts, sink, []*big.Int{})
	}, wsServer, queueService, db)
}
