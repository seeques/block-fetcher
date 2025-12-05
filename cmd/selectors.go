package cmd

import (
	// "context"
	"fmt"
	"log"
	"strings"

	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi"
	contracts "github.com/seeques/block-fetcher/contracts"
	"github.com/spf13/cobra"
)

var inputData []byte

var selectorsCmd = &cobra.Command{
	Use:   "selectors",
	Short: "Decodes transaction input data",
	Long:  `Decodes transaction input data: fetches selectors and function's arguments`,
	Run: func(cmd *cobra.Command, data []string) {
		if len(inputData) < 4 {
			log.Fatalf("Input data is too short to contain a selector")
		}

		contractABI, err := abi.JSON(strings.NewReader(contracts.ContractsMetaData.ABI))
		if err != nil {
			log.Fatalf("Failed to parse contract ABI: %v", err)
		}

		selector := inputData[:4]
		method, err := contractABI.MethodById(selector)
		if err != nil {
			log.Fatalf("Failed to find method for selector %x: %v", selector, err)
		}

		fmt.Printf("Function: %s\n", method.Name)

		argsData := inputData[4:]
		args, err := method.Inputs.Unpack(argsData)
		if err != nil {
			log.Fatalf("Failed to unpack arguments: %v", err)
		}

		for _, arg := range args {
			fmt.Printf("Arg: %v\n", arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(selectorsCmd)
	selectorsCmd.Flags().BytesHexVar(&inputData, "data", nil, "Transaction input data in hex format")
}