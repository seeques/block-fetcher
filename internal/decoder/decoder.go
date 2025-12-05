package decoder

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	contracts "github.com/seeques/block-fetcher/contracts"
)

func ComputeSelector(signature string) []byte {
	return crypto.Keccak256([]byte(signature))[:4]
}

func DecodeData(inputData []byte) (string, []byte, []interface{}, error) {
	contractABI, err := abi.JSON(strings.NewReader(contracts.ContractsMetaData.ABI))
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	selector := inputData[:4]

	method, err := contractABI.MethodById(selector)
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to find method for selector %x: %v", selector, err)
	}

	argsData := inputData[4:]
	args, err := method.Inputs.Unpack(argsData)
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to unpack arguments: %v", err)
	}

	return method.Name, selector, args, nil
}
