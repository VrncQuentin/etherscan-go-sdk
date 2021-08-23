package data

import "math/big"

type (
	Token struct {
		Name    string `json:"tokenName"`
		Symbol  string `json:"tokenSymbol"`
		Decimal string `json:"tokenDecimal"`
	}

	Wei = big.Int
)
