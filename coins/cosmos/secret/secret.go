package secret

import "github.com/mmdi1/go-wallet-sdk/coins/cosmos"

const (
	HRP = "secret"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
