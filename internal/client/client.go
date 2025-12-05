package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectToClient(rpcURL string) (*ethclient.Client) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("failed to connect to the Ethereum client: %v", err)
	}
	return client
}