package faucet

import (
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
)

type FundAccountResult struct {
}

func FundAccount(address string, amount uint64) ([]string, error) {
	client := rpc.New("https://faucet.devnet.aptoslabs.com")
	result, err := client.Post("/mint", map[string]string{
		"amount":  fmt.Sprintf("%d", amount),
		"address": address,
	}, nil)
	if err != nil {
		return nil, err
	}
	var hashs []string
	if err = json.Unmarshal(result, &hashs); err != nil {
		return nil, err
	}
	return hashs, nil
}
