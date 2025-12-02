package main

import (
	"fmt"
	"log"
	"context"
	"math/big"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// Using anvil for now
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection!")

	defer client.Close()

	// Fetch the latest block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latest block number is %d\n", block.Number().Uint64())

	// Let's fetch the unix time and convert it to a more readable format
	timestamp := block.Time()
	fmt.Printf("Block timestamp is %d\n", timestamp)

	t := time.Unix(int64(timestamp), 0)
	formatted := t.Format("02 Jan 2006 15:04:05")
	fmt.Printf("Block creation time is %s\n", formatted)

	// First account from anvil
	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Raw WEI balance is %s\n", balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("ETH balance is %f\n", ethValue)
}