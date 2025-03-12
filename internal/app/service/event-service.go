package service

import (
	"encoding/json"
	"log"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	financialcompliance "monitor-service/internal/modules/compliance/FinancialRWA"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// EventService gerencia a lógica de eventos
func EventService(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, contractAddress string) {
	go wsServer.HandleMessages()

	contract, err := financialcompliance.NewFinancialcompliance(common.HexToAddress(contractAddress), rpcClient.Client())
	if err != nil {
		log.Fatal("Erro ao criar instância do contrato:", err)
	}

	// Events Listeners
	go func() {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceBuyerApproved)
		sub, err := contract.WatchBuyerApproved(opts, eventCh, nil)
		if err != nil {
			log.Fatal("Erro ao escutar eventos:", err)
		}
		defer sub.Unsubscribe()

		for event := range eventCh {
			log.Println("Evento capturado:", event)
			eventBytes, err := json.Marshal(event)
			if err != nil {
				log.Println("Erro ao converter evento para JSON:", err)
				continue
			}
			wsServer.Broadcast(eventBytes)
		}
	}()
}
