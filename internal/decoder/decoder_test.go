package decoder

import (
	"testing"
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func TestDecodeData(t *testing.T) {
	test := struct {
		name string
		inputData []byte
		expectedMethod string
		expectedSel []byte
	}{
		name: "transfer",
		// https://etherscan.io/tx/0xc9535774c0610689c049caf9fa9cef2c3f30a2657d8efc9527ee371619287744
		inputData: common.FromHex("0xa9059cbb0000000000000000000000006bdaa62de230e76cf56662dfa8aa4159362bddd100000000000000000000000000000000000000000000000000027d4afffb7569"),
		expectedMethod: "transfer",
		expectedSel: []byte{0xa9, 0x05, 0x9c, 0xbb},
	}

	methodName, selector, args, err := DecodeData(test.inputData)
	if err != nil {
		t.Fatalf("DecodeData failed: %v", err)
	}

	if methodName != test.expectedMethod {
		t.Errorf("Expected method name %s, got %s", test.expectedMethod, methodName)
	}

	if !bytes.Equal(selector, test.expectedSel) {
		t.Errorf("Expected selector %x, got %x", test.expectedSel, selector)
	}

	// Verify args
	if len(args) != 2 {
		t.Fatalf("Expected 2 arguments, got %d", len(args))
	}

	toAddress, ok := args[0].(common.Address)
	if !ok {
		t.Fatalf("Expected first argument to be common.Address, got %T", args[0])
	}
	expectedTo := common.HexToAddress("6bdaa62de230e76cf56662dfa8aa4159362bddd1")
	if toAddress != expectedTo {
		t.Errorf("Expected to address %s, got %s", expectedTo.Hex(), toAddress.Hex())
	}

	value, ok := args[1].(*big.Int)
	if !ok {
		t.Fatalf("Expected second argument to be *big.Int, got %T", args[1])
	}
	expectedValue := big.NewInt(700711029142889)
	// 0 for equal
	if value.Cmp(expectedValue) != 0 {
		t.Errorf("Expected value %s, got %s", expectedValue.String(), value.String())
	}
}