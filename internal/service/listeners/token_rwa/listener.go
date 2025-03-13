package token_rwa

import (
	"context"
	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/utils"

	token "monitor-service/internal/modules/Token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error creating Token contract instance")
		return
	}

	go listenRealTimeEvents(rpcClient, "Approval", func(opts *bind.WatchOpts, sink chan<- *token.TokenApproval) (event.Subscription, error) {
		return contract.WatchApproval(opts, sink, nil, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "Transfer", func(opts *bind.WatchOpts, sink chan<- *token.TokenTransfer) (event.Subscription, error) {
		return contract.WatchTransfer(opts, sink, nil, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "Paused", contract.WatchPaused, wsServer, queueService, db, "token_rwa")
	go listenRealTimeEvents(rpcClient, "Unpaused", contract.WatchUnpaused, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "OwnershipTransferred", func(opts *bind.WatchOpts, sink chan<- *token.TokenOwnershipTransferred) (event.Subscription, error) {
		return contract.WatchOwnershipTransferred(opts, sink, nil, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "TokensFrozen", func(opts *bind.WatchOpts, sink chan<- *token.TokenTokensFrozen) (event.Subscription, error) {
		return contract.WatchTokensFrozen(opts, sink, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "TokensUnfrozen", func(opts *bind.WatchOpts, sink chan<- *token.TokenTokensUnfrozen) (event.Subscription, error) {
		return contract.WatchTokensUnfrozen(opts, sink, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "RecoverySuccess", func(opts *bind.WatchOpts, sink chan<- *token.TokenRecoverySuccess) (event.Subscription, error) {
		return contract.WatchRecoverySuccess(opts, sink, nil, nil, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "UpdatedTokenInformation", func(opts *bind.WatchOpts, sink chan<- *token.TokenUpdatedTokenInformation) (event.Subscription, error) {
		return contract.WatchUpdatedTokenInformation(opts, sink, nil, nil, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "AddressFrozen", func(opts *bind.WatchOpts, sink chan<- *token.TokenAddressFrozen) (event.Subscription, error) {
		return contract.WatchAddressFrozen(opts, sink, nil, nil, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "AgentAdded", func(opts *bind.WatchOpts, sink chan<- *token.TokenAgentAdded) (event.Subscription, error) {
		return contract.WatchAgentAdded(opts, sink, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "AgentRemoved", func(opts *bind.WatchOpts, sink chan<- *token.TokenAgentRemoved) (event.Subscription, error) {
		return contract.WatchAgentRemoved(opts, sink, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "ComplianceAdded", func(opts *bind.WatchOpts, sink chan<- *token.TokenComplianceAdded) (event.Subscription, error) {
		return contract.WatchComplianceAdded(opts, sink, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "IdentityRegistryAdded", func(opts *bind.WatchOpts, sink chan<- *token.TokenIdentityRegistryAdded) (event.Subscription, error) {
		return contract.WatchIdentityRegistryAdded(opts, sink, nil)
	}, wsServer, queueService, db, "token_rwa")

	go listenRealTimeEvents(rpcClient, "Initialized", func(opts *bind.WatchOpts, sink chan<- *token.TokenInitialized) (event.Subscription, error) {
		return contract.WatchInitialized(opts, sink)
	}, wsServer, queueService, db, "token_rwa")

	logger.Log.Info().Msg("🎯 WebSocket event listeners started successfully")
}

func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*token.TokenFilterer, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return token.NewTokenFilterer(common.HexToAddress(contractAddress), client)
}

func listenRealTimeEvents[T any](rpcClient *rpc.RPCClient, eventName string, watchFunc func(*bind.WatchOpts, chan<- T) (event.Subscription, error), wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, collectionName string) {
	sink := make(chan T)

	client, _ := rpcClient.Client()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error getting latest block")
		return
	}

	headerNumber := header.Number.Uint64()
	opts := &bind.WatchOpts{
		Start:   &headerNumber,
		Context: context.Background(),
	}

	sub, err := watchFunc(opts, sink)
	if err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error starting event listener")
		return
	}
	defer sub.Unsubscribe()

	logger.Log.Info().Str("event", eventName).Msg("📡 WebSocket listener started")

	for event := range sink {
		utils.ProcessEvent(eventName, event, wsServer, queueService, db, collectionName)
	}
}
