package financial_compliance

import (
	"encoding/json"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	financialcompliance "monitor-service/internal/modules/compliance/FinancialRWA"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func StartListener(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, contractAddress string) {
	contract, err := initializeContract(rpcClient, contractAddress)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Error creating FinancialCompliance contract instance")
		return
	}

	go listenForBuyerApproved(contract, wsServer, queueService, db)
	go listenForComplianceBound(contract, wsServer, queueService, db)
	go listenForComplianceCheckFailed(contract, wsServer, queueService, db)
	go listenForComplianceCheckPassed(contract, wsServer, queueService, db)
	go listenForComplianceUnbound(contract, wsServer, queueService, db)
	go listenForOwnershipTransferred(contract, wsServer, queueService, db)
}

func initializeContract(rpcClient *rpc.RPCClient, contractAddress string) (*financialcompliance.Financialcompliance, error) {
	client, err := rpcClient.Client()
	if err != nil {
		return nil, err
	}
	return financialcompliance.NewFinancialcompliance(common.HexToAddress(contractAddress), client)
}

func listenForBuyerApproved(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceBuyerApproved)

		sub, err := contract.WatchBuyerApproved(opts, eventCh, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to BuyerApproved. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("BuyerApproved", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForComplianceBound(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceComplianceBound)

		sub, err := contract.WatchComplianceBound(opts, eventCh, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to ComplianceBound. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("ComplianceBound", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForComplianceCheckFailed(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceComplianceCheckFailedEvent)

		sub, err := contract.WatchComplianceCheckFailedEvent(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to ComplianceCheckFailed. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("ComplianceCheckFailed", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForComplianceCheckPassed(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceComplianceCheckPassed)

		sub, err := contract.WatchComplianceCheckPassed(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to ComplianceCheckPassed. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("ComplianceCheckPassed", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForComplianceUnbound(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceComplianceUnbound)

		sub, err := contract.WatchComplianceUnbound(opts, eventCh, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to ComplianceUnbound. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("ComplianceUnbound", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
	}
}

func listenForOwnershipTransferred(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceOwnershipTransferred)

		sub, err := contract.WatchOwnershipTransferred(opts, eventCh, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to OwnershipTransferred. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			processEvent("OwnershipTransferred", event, wsServer, queueService, db)
		}

		sub.Unsubscribe()
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
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error decoding event JSON")
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
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error formatting event JSON")
		return
	}
	logger.Log.Info().Str("event", eventName).RawJSON("formatted_data", formattedBytes).Msg("📢 Event captured")

	wsServer.Broadcast(formattedBytes)

	if err := queueService.Publish(formattedBytes); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error publishing event to RabbitMQ")
	}

	if err := db.InsertEvent("financial_compliance", formattedEvent); err != nil {
		logger.Log.Error().Err(err).Str("event", eventName).Msg("❌ Error saving event to MongoDB")
	}
}
