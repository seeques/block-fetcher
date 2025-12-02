package main

import (
	"testing"
	// "context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func setUpConnection(t *testing.T) *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to client: %v", err)
	}
	return client
}

func TestConnection(t *testing.T) {
	client := setUpConnection(t)
	defer client.Close()
}