package service

import (
	"encoding/json"
	"log"
	"time"

	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	financialcompliance "monitor-service/internal/modules/compliance/FinancialRWA"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// EventService gerencia a lógica de eventos do contrato
func EventService(rpcClient *rpc.RPCClient, wsServer *websocket.WebSocketServer, contractAddress string) {
	go wsServer.HandleMessages()

	// Inicia loop para garantir conexão válida com o contrato
	var contract *financialcompliance.Financialcompliance
	var err error
	for i := 0; i < 5; i++ {
		client, err := rpcClient.Client()
		if err == nil {
			contract, err = financialcompliance.NewFinancialcompliance(common.HexToAddress(contractAddress), client)
			if err == nil {
				break
			}
		}

		log.Printf("⚠️ Tentando novamente inicializar contrato... (%d/5)\n", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("❌ Erro ao criar instância do contrato após tentativas: %v", err)
	}

	// Inicia monitoramento de eventos em uma goroutine separada
	go listenForBuyerApprovedEvents(contract, wsServer)
}

// listenForBuyerApprovedEvents escuta eventos BuyerApproved com reconexão automática
func listenForBuyerApprovedEvents(contract *financialcompliance.Financialcompliance, wsServer *websocket.WebSocketServer) {
	for {
		opts := &bind.WatchOpts{}
		eventCh := make(chan *financialcompliance.FinancialcomplianceBuyerApproved)

		sub, err := contract.WatchBuyerApproved(opts, eventCh, nil)
		if err != nil {
			log.Printf("❌ Erro ao escutar eventos: %v. Tentando novamente em 5s...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer sub.Unsubscribe()

		// Loop para capturar eventos
		for event := range eventCh {
			log.Println("📢 Evento capturado:", event)
			eventBytes, err := json.Marshal(event)
			if err != nil {
				log.Println("❌ Erro ao converter evento para JSON:", err)
				continue
			}
			wsServer.Broadcast(eventBytes)
		}

		log.Println("⚠️ Canal de eventos fechado. Tentando reiniciar escuta...")
		time.Sleep(3 * time.Second)
	}
}
