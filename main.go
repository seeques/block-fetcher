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

	// Fetch all txs in a block
	for _, tx := range block.Transactions() {
		fmt.Printf("Tx hash is: %s\n", tx.Hash().Hex())

		txValuef := new(big.Float).SetInt(tx.Value())
		gasPricef := new(big.Float).SetInt(tx.GasPrice())
		fmt.Printf("Tx value in ETH is: %s\n", new(big.Float).Quo(txValuef, big.NewFloat(math.Pow10(18))).Text('f', 18))
		fmt.Printf("Tx gas price in ETH is: %s\n", new(big.Float).Quo(gasPricef, big.NewFloat(math.Pow10(18))).Text('f', 18))
	}
}