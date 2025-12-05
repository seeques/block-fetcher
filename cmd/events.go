package cmd

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	contracts "github.com/seeques/block-fetcher/contracts"
	"github.com/seeques/block-fetcher/internal/client"
	"github.com/seeques/block-fetcher/internal/fetcher"
	"github.com/spf13/cobra"
)

var contractAddress string
var block int64

type TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Fetches ERC20 transfer event from one block",
	Long:  `Fetches ERC20 transfer event from a specified contract address from one block, latest by default`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.ConnectToClient(rpcURL)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		b, err := fetcher.GetBlockNumber(client, block)
		if err != nil {
			log.Fatal(err)
		}

		logs, err := fetcher.GetLogs(client, contractAddress, b)
		if err != nil {
			log.Fatal(err)
		}

		contractABI, err := abi.JSON(strings.NewReader(string(contracts.ContractsMetaData.ABI)))
		if err != nil {
			log.Fatalf("Failed to parse contract ABI: %v", err)
		}

		logTransferSig := []byte("Transfer(address,address,uint256)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

		for _, vLog := range logs {
			if vLog.Topics[0] == logTransferSigHash {
				fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
				fmt.Printf("Log Index: %d\n", vLog.Index)

				fmt.Printf("Log name: %s\n", "Transfer")

				var transferEvent TransferEvent
				// Unpack can only unpack non-indexed fields of event
				// Both from and to are indexed fields that sit in topics
				err := contractABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatalf("Failed to unpack log: %v", err)
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Value: %s\n", transferEvent.Value.String())
				fmt.Println("-----")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.Flags().StringVarP(&contractAddress, "address", "a", "", "ERC20 contract address to fetch transfer events from")
	eventsCmd.Flags().Int64VarP(&block, "block", "b", 0, "Block number to fetch events from")
}
