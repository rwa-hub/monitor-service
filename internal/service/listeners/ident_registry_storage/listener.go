package ident_registry_storage

import (
	"encoding/json"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	identregistrystorage "monitor-service/internal/modules/IdentRegistryStorage"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Erro ao criar instância do contrato IdentRegistryStorage")
		return
	}

	go listenForIdentityStored(contract, wsServer, queueService, db)
	go listenForIdentityModified(contract, wsServer, queueService, db)
	go listenForIdentityUnstored(contract, wsServer, queueService, db)
}

func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*identregistrystorage.Identregistrystorage, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return identregistrystorage.NewIdentregistrystorage(common.HexToAddress(contractAddress), client)
}

func processEvent(eventType string, event interface{}, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao converter evento para JSON")
		return
	}

	wsServer.Broadcast(eventBytes)

	err = queueService.Publish(eventBytes)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao enviar evento para RabbitMQ")
	}

	err = db.InsertEvent(eventType, event)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao armazenar evento no MongoDB")
	}

	logger.Log.Info().Str("eventType", eventType).Msg("📢 Evento processado com sucesso")
}

func listenForIdentityStored(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityStored)

		sub, err := contract.WatchIdentityStored(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Erro ao escutar eventos IdentityStored. Tentando novamente em 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("IdentityStored", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForIdentityModified(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityModified)

		sub, err := contract.WatchIdentityModified(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Erro ao escutar eventos IdentityModified. Tentando novamente em 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("IdentityModified", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForIdentityUnstored(contract *identregistrystorage.Identregistrystorage, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *identregistrystorage.IdentregistrystorageIdentityUnstored)

		sub, err := contract.WatchIdentityUnstored(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Erro ao escutar eventos IdentityUnstored. Tentando novamente em 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("IdentityUnstored", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}
