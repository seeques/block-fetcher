package cmd

import (
	"fmt"
	"log"
	"bytes"

	"github.com/seeques/block-fetcher/internal/decoder"
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

		methodName, selector, args, err := decoder.DecodeData(inputData)
		if err != nil {
			log.Fatal(err)
		}

		switch {
			case bytes.Equal(selector, decoder.ComputeSelector("transfer(address,uint256)")):
				fmt.Printf("Method: %s\n", methodName)
				fmt.Printf("To: %v\n", args[0])
				fmt.Printf("Value: %v\n", args[1])
			case bytes.Equal(selector, decoder.ComputeSelector("approve(address,uint256)")):
				fmt.Printf("Method: %s\n", methodName)
				fmt.Printf("Spender: %v\n", args[0])
				fmt.Printf("Value: %v\n", args[1])
			case bytes.Equal(selector, decoder.ComputeSelector("transferFrom(address,address,uint256)")):
				fmt.Printf("Method: %s\n", methodName)
				fmt.Printf("From: %v\n", args[0])
				fmt.Printf("To: %v\n", args[1])
				fmt.Printf("Value: %v\n", args[2])
			default:
				fmt.Printf("Method: %s\n", methodName)
				fmt.Println("Arguments:")
				for i, arg := range args {
					fmt.Printf("Arg %d: %v\n", i, arg)
				}
		}
	},
}

func init() {
	rootCmd.AddCommand(selectorsCmd)
	selectorsCmd.Flags().BytesHexVar(&inputData, "data", nil, "Transaction input data in hex format")
}