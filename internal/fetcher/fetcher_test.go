package fetcher

import (
	"testing"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func setUpConnection(t *testing.T) *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to client: %v", err)
	}
	return client
}

func TestGetBlockNumber(t *testing.T) {
	client := setUpConnection(t)
	defer client.Close()
	
	blockNum, err := GetBlockNumber(client, 0)
	if err != nil {
		t.Fatalf("Failed to get block number: %v", err)
	}
	if blockNum.Cmp(big.NewInt(0)) == -1 {
		t.Fatalf("Block number should not be negative")
	}
}