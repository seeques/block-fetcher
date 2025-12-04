package main

import (
	"context"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func TestFetchBlock(t *testing.T) {
	client := setUpConnection(t)
	defer client.Close()

	_, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to fetch latest block: %v", err)
	}
}

func TestReadEthValue(t *testing.T) {
	client := setUpConnection(t)
	defer client.Close()

	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		t.Fatalf("Failed to fetch account balance: %v", err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// cmp returns -1 on negative
	if ethValue.Cmp(big.NewFloat(0)) == -1 {
		t.Fatalf("ETH balance should not be negative")
	}
}

func TestFetchTxValueNonNegative(t *testing.T) {
	client := setUpConnection(t)
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to fetch latest block: %v", err)
	}

	txs := block.Transactions()
	if len(txs) == 0 {
		t.Skip("No txs in the latest block to test")
	}

	tx := txs[0]
	value := tx.Value()

	if value.Sign() == -1 {
		t.Fatalf("Transaction value should not be negative")
	}
}
