package financial_compliance

import (
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	financialcompliance "monitor-service/internal/modules/compliance/FinancialRWA"
	"monitor-service/utils"

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

		sub, err := contract.WatchBuyerApproved(opts, eventCh, nil, nil, nil)
		if err != nil {
			logger.Log.Error().Err(err).Msg("❌ Error listening to BuyerApproved. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range eventCh {
			utils.ProcessEvent("BuyerApproved", event, wsServer, queueService, db, "financial_compliance")
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
			utils.ProcessEvent("ComplianceBound", event, wsServer, queueService, db, "financial_compliance")
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
			utils.ProcessEvent("ComplianceCheckFailed", event, wsServer, queueService, db, "financial_compliance")
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
			utils.ProcessEvent("ComplianceCheckPassed", event, wsServer, queueService, db, "financial_compliance")
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
			utils.ProcessEvent("ComplianceUnbound", event, wsServer, queueService, db, "financial_compliance")
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
			utils.ProcessEvent("OwnershipTransferred", event, wsServer, queueService, db, "financial_compliance")
		}

		sub.Unsubscribe()
	}
}
