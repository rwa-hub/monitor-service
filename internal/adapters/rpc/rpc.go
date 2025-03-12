package rpc

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RPCClient gerencia a conexão WebSocket e eventos do Ethereum
type RPCClient struct {
	client    *ethclient.Client
	rpcURL    string
	events    chan types.Log
	mu        sync.Mutex
	connected bool
}

// NewRPCClient cria uma nova instância de RPCClient
func NewRPCClient(rpcURL string) (*RPCClient, error) {
	rpc := &RPCClient{
		rpcURL:    rpcURL,
		events:    make(chan types.Log, 100),
		connected: false,
	}

	if err := rpc.connect(); err != nil {
		return nil, err
	}

	return rpc, nil
}

// connect inicia a conexão com o WebSocket do Ethereum
func (r *RPCClient) connect() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	client, err := ethclient.Dial(r.rpcURL)
	if err != nil {
		log.Printf("❌ Falha ao conectar ao nó Ethereum: %v", err)
		return err
	}

	r.client = client
	r.connected = true
	log.Println("✅ Conectado ao Ethereum WebSocket com sucesso!")
	return nil
}

// Client retorna a instância do cliente Ethereum, garantindo que nunca seja nil
func (r *RPCClient) Client() (*ethclient.Client, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.client == nil {
		log.Println("⚠️ Cliente Ethereum está nil. Tentando reconectar...")
		if err := r.connect(); err != nil {
			return nil, err
		}
	}

	return r.client, nil
}

// Reconnect força a reconexão ao nó Ethereum
func (r *RPCClient) Reconnect() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	log.Println("⚠️ Tentando reconectar ao nó Ethereum WebSocket...")
	for i := 1; i <= 5; i++ {
		client, err := ethclient.Dial(r.rpcURL)
		if err == nil {
			r.client = client
			r.connected = true
			log.Println("✅ Reconectado com sucesso ao Ethereum WebSocket!")
			return nil
		}

		log.Printf("❌ Falha ao reconectar (%d/5). Tentando novamente em 3s...\n", i)
		time.Sleep(3 * time.Second)
	}

	return errors.New("falha ao reconectar após várias tentativas")
}

// ListenEvents escuta eventos de contratos e reconecta caso a conexão caia
func (r *RPCClient) ListenEvents(contractAddress string, eventSignature string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	for {
		client, err := r.Client()
		if err != nil {
			log.Println("⚠️ Cliente Ethereum ainda indisponível. Tentando reconectar...")
			time.Sleep(3 * time.Second)
			continue
		}

		log.Println("🎧 Iniciando escuta de eventos...")

		logs := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Printf("❌ Erro ao iniciar a subscrição: %v\n", err)
			time.Sleep(3 * time.Second)
			continue
		}

		for {
			select {
			case err := <-sub.Err():
				log.Printf("❌ Erro na conexão WebSocket: %v\n", err)
				sub.Unsubscribe()
				r.connected = false
				time.Sleep(3 * time.Second)
				break
			case vLog := <-logs:
				log.Println("📢 Evento capturado:", vLog)
				r.events <- vLog
			}
		}
	}
}
