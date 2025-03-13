package identity_registry

import (
	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	identityregistry "monitor-service/internal/modules/IdentityRegistry"
	"monitor-service/utils"

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
		utils.ProcessEvent("IdentityRegistered", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("IdentityRemoved", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("IdentityUpdated", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("ClaimTopicsRegistrySet", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("TrustedIssuersRegistrySet", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("IdentityStorageSet", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("CountryUpdated", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("AgentAdded", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("AgentRemoved", event, wsServer, queueService, db, "identity_registry")
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
		utils.ProcessEvent("OwnershipTransferred", event, wsServer, queueService, db, "identity_registry")
	}
}
