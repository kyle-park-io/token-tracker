package call

import (
	"math/big"
	"testing"
)

// go test -v -run TestMakeCallData
func TestMakeCallData(t *testing.T) {

	// Example: ERC-20 balanceOf(address)
	methodName := "balanceOf"
	paramTypes := []string{"address"}
	params := []interface{}{"0x5c7BCd6E7De5423a257D81B442095A1a6ced35C5"}

	callData, err := CreateCallData(methodName, paramTypes, params)
	if err != nil {
		t.Error("Error:", err)
		return
	}
	t.Log("Call Data:", callData)

	// Example: transfer(address, uint256)
	methodName = "transfer"
	paramTypes = []string{"address", "uint256"}
	params = []interface{}{
		"0x5c7BCd6E7De5423a257D81B442095A1a6ced35C5",
		big.NewInt(1000000000000000000), // 1 ETH (in wei)
	}

	callData, err = CreateCallData(methodName, paramTypes, params)
	if err != nil {
		t.Error("Error:", err)
		return
	}
	t.Log("Call Data:", callData)

}
