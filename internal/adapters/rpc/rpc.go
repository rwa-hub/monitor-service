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

type RPCClient struct {
	client    *ethclient.Client
	rpcURL    string
	events    chan types.Log
	mu        sync.Mutex
	connected bool
}

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

func (r *RPCClient) connect() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	client, err := ethclient.Dial(r.rpcURL)
	if err != nil {
		log.Printf("❌ Error to connect to Ethereum node: %v", err)
		return err
	}

	r.client = client
	r.connected = true
	log.Println("--> Connected to Ethereum WebSocket successfully!")
	return nil
}

func (r *RPCClient) Client() (*ethclient.Client, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.client == nil {
		log.Println("--> Ethereum client is nil. Trying to reconnect...")
		if err := r.connect(); err != nil {
			return nil, err
		}
	}

	return r.client, nil
}

func (r *RPCClient) Reconnect() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	log.Println("--> Trying to reconnect to Ethereum WebSocket...")
	for i := 1; i <= 5; i++ {
		client, err := ethclient.Dial(r.rpcURL)
		if err == nil {
			r.client = client
			r.connected = true
			log.Println("--> Reconnected successfully to Ethereum WebSocket!")
			return nil
		}

		log.Printf("❌ Error to reconnect (%d/5). Retrying in 3s...\n", i)
		time.Sleep(3 * time.Second)
	}

	return errors.New("error to reconnect after several attempts")
}

func (r *RPCClient) ListenEvents(contractAddress string, eventSignature string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	for {
		client, err := r.Client()
		if err != nil {
			log.Println("--> Ethereum client is still unavailable. Trying to reconnect...")
			time.Sleep(3 * time.Second)
			continue
		}

		log.Println("--> Starting event listening...")

		logs := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Printf("❌ Error to start subscription: %v\n", err)
			time.Sleep(3 * time.Second)
			continue
		}

	outerLoop:
		for {
			select {
			case err := <-sub.Err():
				log.Printf("❌ Error in WebSocket connection: %v\n", err)
				sub.Unsubscribe()
				r.connected = false
				time.Sleep(3 * time.Second)
				break outerLoop
			case vLog := <-logs:
				log.Println("--> Event captured:", vLog)
				r.events <- vLog
			}
		}
	}
}
