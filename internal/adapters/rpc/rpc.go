package rpc

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RPCClient é responsável por gerenciar a conexão com o nó Ethereum
type RPCClient struct {
	client *ethclient.Client
	events chan types.Log
}

// NewRPCClient cria uma nova instância de RPCClient
func NewRPCClient(rpcURL string) (*RPCClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &RPCClient{client: client, events: make(chan types.Log)}, nil
}

// EventChannel retorna o canal de eventos
func (r *RPCClient) EventChannel() <-chan types.Log {
	return r.events
}

// ListenEvents escuta eventos de contratos inteligentes
func (r *RPCClient) ListenEvents(contractAddress string, eventSignature string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	logs := make(chan types.Log)
	sub, err := r.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case vLog := <-logs:
			log.Println("Evento capturado:", vLog)
			r.events <- vLog // Enviar evento para o canal
		}
	}
}
