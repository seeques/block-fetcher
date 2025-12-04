package cmd

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	// "time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var blockNumber int64
var showReceipts bool

var txsCmd = &cobra.Command{
	Use: "txs",
	Short: "Fetches transaction data from block",
	Long: `Fetches transaction data from a specified block or the latest block as default`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		defer client.Close()

		fmt.Println("Connected to:", rpcURL)
		fmt.Println("-----")

		var block *types.Block 

		if blockNumber < 0 {
			// nil for latest
			block, err = client.BlockByNumber(context.Background(), nil)
		} else {
			block, err = client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
		}
		if err != nil {
			log.Fatalf("Failed to fetch block: %v", err)
		}
		
		fmt.Printf("Block Number: %d\n", block.Number().Uint64())
		fmt.Println("-----")
		
		// We need chainID to get the signer
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatalf("Failed to get network ID: %v", err)
		}

		// Fetch all txs in a block with their respective data
		for _, tx := range block.Transactions() {
			fmt.Printf("Tx hash: %s\n", tx.Hash().Hex())

			txValuef := new(big.Float).SetInt(tx.Value())
			ethValue := new(big.Float).Quo(txValuef, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Value: %s Eth\n", ethValue.Text('f', 18))

			gasPricef := new(big.Float).SetInt(tx.GasPrice())
			gasPriceEth := new(big.Float).Quo(gasPricef, big.NewFloat(math.Pow10(18)))
			fmt.Printf("Gas Price: %s Eth\n", gasPriceEth.Text('f', 18))

			if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
				fmt.Printf("From: %s\n", from.Hex())
			}

			to := tx.To()
			if to != nil {
				fmt.Printf("To: %s\n", to.Hex())
			} else {
				fmt.Println("To: Contract Creation")
			}
			
			fmt.Printf("Gas Limit: %d\n", tx.Gas())
			fmt.Printf("Nonce: %d\n", tx.Nonce())
			
			txData := tx.Data()
			if len(txData) > 0 {
				fmt.Printf("Data: 0x%x\n", txData)
			} else {
				fmt.Println("Data: <empty>")
			}

			// Show receipts if flag is set
			if showReceipts {
				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Printf("Failed to fetch receipt for tx %s: %v", tx.Hash().Hex(), err)
				} else {
					fmt.Printf("Receipt Status: %d\n", receipt.Status)
					// TODO: think about better formatting for logs
					fmt.Printf("Logs: %v\n", receipt.Logs)
				}
			}
			fmt.Println("-----")
		}
	},
}

func init() {
	rootCmd.AddCommand(txsCmd)
	txsCmd.Flags().Int64VarP(&blockNumber, "block", "b", -1, "Block Number to fetch, -1 for latest")
	txsCmd.Flags().BoolVarP(&showReceipts, "receipts", "r", false, "Show transaction receipts")
}