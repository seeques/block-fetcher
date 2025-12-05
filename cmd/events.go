package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/seeques/block-fetcher/contracts"
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
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		defer client.Close()

		if contractAddress == "" {
			log.Fatalf("Contract address is required")
		}

		contractAddr := common.HexToAddress(contractAddress)

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

		query := ethereum.FilterQuery{
			// Topics == nil (default) matches all topics
			Addresses: []common.Address{
				contractAddr,
			},
			FromBlock: b,
			ToBlock:   b,
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatalf("Failed to retrieve logs: %v", err)
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
