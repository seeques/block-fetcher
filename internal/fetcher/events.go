package fetcher

import (
	"context"
	// "fmt"
	"log"
	"math/big"
	// "strings"

	// "github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// contracts "github.com/seeques/block-fetcher/contracts"
)

func GetBlock(client *ethclient.Client, block int64) *big.Int {
	var b *big.Int
	// default to latest block if 0
	if block == 0 {
		// Get latest block number
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to get latest block header: %v", err)
		}
		b = header.Number
	} else {
		b = big.NewInt(block)
	}
	return b
}
