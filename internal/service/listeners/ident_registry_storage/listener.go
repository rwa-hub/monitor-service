package ident_registry_storage

import (
	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	identregistrystorage "monitor-service/internal/modules/IdentRegistryStorage"
	"monitor-service/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Definir uma interface para o iterador de eventos
type EventIterator interface {
	Next() bool
	Event() interface{}
}

// StartListener inicia os listeners para todos os eventos do contrato IdentRegistryStorage
func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error creating IdentRegistryStorage contract instance")
		return
	}

	// Lista de todos os eventos suportados
	eventHandlers := map[string]func(){
		"AgentAdded":              func() { watchAgentAdded(contract, wsServer, queueService, db) },
		"AgentRemoved":            func() { watchAgentRemoved(contract, wsServer, queueService, db) },
		"CountryModified":         func() { watchCountryModified(contract, wsServer, queueService, db) },
		"IdentityModified":        func() { watchIdentityModified(contract, wsServer, queueService, db) },
		"IdentityRegistryBound":   func() { watchIdentityRegistryBound(contract, wsServer, queueService, db) },
		"IdentityRegistryUnbound": func() { watchIdentityRegistryUnbound(contract, wsServer, queueService, db) },
		"IdentityStored":          func() { watchIdentityStored(contract, wsServer, queueService, db) },
		"IdentityUnstored":        func() { watchIdentityUnstored(contract, wsServer, queueService, db) },
		"Initialized":             func() { watchInitialized(contract, wsServer, queueService, db) },
		"OwnershipTransferred":    func() { watchOwnershipTransferred(contract, wsServer, queueService, db) },
	}

	// Inicia um goroutine para cada evento
	for eventName, handler := range eventHandlers {
		go func(e string, h func()) {
			for {
				h()
				logger.Log.Warn().Str("event", e).Msg("🔄 Restarting listener due to error or disconnection...")
				time.Sleep(5 * time.Second)
			}
		}(eventName, handler)
	}
}

// Inicializa o contrato na blockchain
func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*identregistrystorage.Identregistrystorage, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return identregistrystorage.NewIdentregistrystorage(common.HexToAddress(contractAddress), client)
}

// 🟢 Implementação dos watchers para cada evento:

func watchAgentAdded(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageAgentAdded)

	sub, err := contract.WatchAgentAdded(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to AgentAdded")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("AgentAdded", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchAgentRemoved(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageAgentRemoved)

	sub, err := contract.WatchAgentRemoved(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to AgentRemoved")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("AgentRemoved", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchCountryModified(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageCountryModified)

	sub, err := contract.WatchCountryModified(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to CountryModified")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("CountryModified", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchIdentityModified(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityModified)

	sub, err := contract.WatchIdentityModified(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityModified")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("IdentityModified", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchIdentityRegistryBound(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityRegistryBound)

	sub, err := contract.WatchIdentityRegistryBound(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityRegistryBound")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("IdentityRegistryBound", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchIdentityRegistryUnbound(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityRegistryUnbound)

	sub, err := contract.WatchIdentityRegistryUnbound(opts, eventCh, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityRegistryUnbound")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("IdentityRegistryUnbound", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchInitialized(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageInitialized)

	sub, err := contract.WatchInitialized(opts, eventCh)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to Initialized")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("Initialized", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchOwnershipTransferred(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageOwnershipTransferred)

	sub, err := contract.WatchOwnershipTransferred(opts, eventCh, nil, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to OwnershipTransferred")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("OwnershipTransferred", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchIdentityStored(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityStored)

	sub, err := contract.WatchIdentityStored(opts, eventCh, []common.Address{}, []common.Address{})
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityStored")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("IdentityStored", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

func watchIdentityUnstored(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	opts := &bind.WatchOpts{}
	eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityUnstored)

	sub, err := contract.WatchIdentityUnstored(opts, eventCh, []common.Address{}, []common.Address{})
	if err != nil {
		logger.Log.Error().Err(err).Msg("❌ Error listening to IdentityUnstored")
		return
	}

	for event := range eventCh {
		utils.ProcessEvent("IdentityUnstored", event, wsServer, queueService, db, "ident_registry_storage")
	}

	sub.Unsubscribe()
}

// Ajustar a função handleEvent para usar a interface EventIterator
func handleEvent[T EventIterator](iterator T, eventType, collectionName string, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for iterator.Next() {
		utils.ProcessEvent(eventType, iterator.Event(), wsServer, queueService, db, collectionName)
	}
}
