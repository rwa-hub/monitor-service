package identity_registry

import (
	"encoding/json"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	identityregistry "monitor-service/internal/modules/IdentityRegistry"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error initializing IdentityRegistry contract")
		return
	}

	go listenIdentityRegistered(contract, wsServer, queueService, db)
	go listenIdentityRemoved(contract, wsServer, queueService, db)
	go listenIdentityUpdated(contract, wsServer, queueService, db)
	go listenClaimTopicsRegistrySet(contract, wsServer, queueService, db)
	go listenTrustedIssuersRegistrySet(contract, wsServer, queueService, db)
	go listenIdentityStorageSet(contract, wsServer, queueService, db)
	go listenCountryUpdated(contract, wsServer, queueService, db)
	go listenAgentAdded(contract, wsServer, queueService, db)
	go listenAgentRemoved(contract, wsServer, queueService, db)
	go listenOwnershipTransferred(contract, wsServer, queueService, db)
}

func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*identityregistry.Identityregistry, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return identityregistry.NewIdentityregistry(common.HexToAddress(contractAddress), client)
}

func listenIdentityRegistered(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryIdentityRegistered)
	sub, err := contract.WatchIdentityRegistered(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityRegistered")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("IdentityRegistered", event, wsServer, queueService, db)
	}
}

func listenIdentityRemoved(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryIdentityRemoved)
	sub, err := contract.WatchIdentityRemoved(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityRemoved")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("IdentityRemoved", event, wsServer, queueService, db)
	}
}

func listenIdentityUpdated(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryIdentityUpdated)
	sub, err := contract.WatchIdentityUpdated(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityUpdated")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("IdentityUpdated", event, wsServer, queueService, db)
	}
}

func listenClaimTopicsRegistrySet(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryClaimTopicsRegistrySet)
	sub, err := contract.WatchClaimTopicsRegistrySet(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to ClaimTopicsRegistrySet")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("ClaimTopicsRegistrySet", event, wsServer, queueService, db)
	}
}

func listenTrustedIssuersRegistrySet(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryTrustedIssuersRegistrySet)
	sub, err := contract.WatchTrustedIssuersRegistrySet(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to TrustedIssuersRegistrySet")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("TrustedIssuersRegistrySet", event, wsServer, queueService, db)
	}
}

func listenIdentityStorageSet(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryIdentityStorageSet)
	sub, err := contract.WatchIdentityStorageSet(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityStorageSet")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("IdentityStorageSet", event, wsServer, queueService, db)
	}
}

func listenCountryUpdated(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryCountryUpdated)
	sub, err := contract.WatchCountryUpdated(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to CountryUpdated")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("CountryUpdated", event, wsServer, queueService, db)
	}
}

func listenAgentAdded(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryAgentAdded)
	sub, err := contract.WatchAgentAdded(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to AgentAdded")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("AgentAdded", event, wsServer, queueService, db)
	}
}

func listenAgentRemoved(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryAgentRemoved)
	sub, err := contract.WatchAgentRemoved(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to AgentRemoved")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("AgentRemoved", event, wsServer, queueService, db)
	}
}

func listenOwnershipTransferred(contract *identityregistry.Identityregistry, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identityregistry.IdentityregistryOwnershipTransferred)
	sub, err := contract.WatchOwnershipTransferred(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to OwnershipTransferred")
		return
	}
	defer sub.Unsubscribe()

	for event := range eventCh {
		processEvent("OwnershipTransferred", event, wsServer, queueService, db)
	}
}

func processEvent(eventName string, event interface{}, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {

	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error converting event to JSON")
		return
	}

	var rawData map[string]interface{}
	if err := json.Unmarshal(eventBytes, &rawData); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error parsing event data")
		return
	}

	formattedEvent := map[string]interface{}{
		"eventName":       eventName,
		"timestamp":       time.Now().UTC().Format(time.RFC3339),
		"contractAddress": rawData["Raw"].(map[string]interface{})["address"],
		"transactionHash": rawData["Raw"].(map[string]interface{})["transactionHash"],
		"blockNumber":     rawData["Raw"].(map[string]interface{})["blockNumber"],
		"data":            rawData,
	}

	formattedBytes, err := json.Marshal(formattedEvent)
	if err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error formatting event")
		return
	}

	logger.Log.Info().Str("event", eventName).RawJSON("formatted_data", formattedBytes).Msg("📢 Event processed")

	wsServer.Broadcast(formattedBytes)

	if err := queueService.Publish(formattedBytes); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error publishing event to queue")
	}

	if err := db.InsertEvent("events", formattedEvent); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error saving event to MongoDB")
	}
}
