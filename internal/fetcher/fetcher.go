package fetcher

import (
	"context"
	"fmt"
	// "log"
	"math/big"
	// "strings"

	"github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	// contracts "github.com/seeques/block-fetcher/contracts"
)

func GetBlockNumber(client *ethclient.Client, block int64) (*big.Int, error) {
	var b *big.Int
	// default to latest block if 0
	if block == 0 {
		// Get latest block number
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to get latest block header: %w", err)
		}
		b = header.Number
	} else {
		b = big.NewInt(block)
	}
	return b, nil
}

func GetLogs(client *ethclient.Client, contractAddress string, block *big.Int) ([]types.Log, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("Contract address is required")
	}

	contractAddr := common.HexToAddress(contractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddr,
		},
		FromBlock: block,
		ToBlock:   block,
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve logs: %w", err)
	}
	return logs, nil
}

func GetBlock(client *ethclient.Client, blockNumber int64) (*types.Block, error) {
	var block *types.Block
	var err error

	if blockNumber < 0 {
		// nil for latest
		block, err = client.BlockByNumber(context.Background(), nil)
	} else {
		block, err = client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch block: %w", err)
	}
	return block, nil
}