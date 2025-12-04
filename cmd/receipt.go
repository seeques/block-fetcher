package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var receiptCmd = &cobra.Command{
	Use:   "receipt",
	Short: "Fetches receipt from transaction hash",
	Long:  `Fetches receipt data from a specified transaction hash`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		defer client.Close()

		if len(args) != 1 {
			log.Fatalf("Transaction hash is required")
		}

		txHash := args[0]

		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Fatalf("Failed to fetch receipt for tx %s: %v", txHash, err)
		}

		fmt.Printf("Logs: %v\n", receipt.Logs)
		fmt.Printf("Receipt Status: %d\n", receipt.Status)
		fmt.Printf("Cumulative Gas Used: %d\n", receipt.CumulativeGasUsed)
		fmt.Printf("EffectiveGasPrice: %d\n", receipt.EffectiveGasPrice)
	},
}

func init() {
	rootCmd.AddCommand(receiptCmd)
	receiptCmd.Flags().String("txhash", "", "Transaction hash to fetch receipt for")
}
