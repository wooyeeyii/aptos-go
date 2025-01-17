package base_example

import (
	"fmt"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestNewAccount(t *testing.T) {
	// new account
	wallet := wallet.New()
	wallet.Save("account_example")
	address := wallet.Address()
	key := wallet.PrivateKey.String()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("key: %s\n", key)
}

func TestReadAccount(t *testing.T) {
	// read account
	wallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	key := wallet.PrivateKey.String()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("key: %s\n", key)
}
