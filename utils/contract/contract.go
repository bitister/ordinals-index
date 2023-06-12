package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func GetAbi(abiStr string) (abi.ABI, error) {
	r := strings.NewReader(abiStr)
	a, err := abi.JSON(r)
	return a, err
}
