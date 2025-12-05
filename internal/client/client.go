package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectToClient(rpcURL string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}
	return client, nil
}