package contracts

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	ABI struct {
		*base.Call
	}

	ABIResult struct {
		base.CallResult
		Data ABIData `json:"result"`
	}

	ABIData struct {
		//TODO: Add data
	}
)

func NewABI(token string) *ABI {
	tx := &ABI{
		base.NewCall(token, new(ABIResult)),
	}
	tx.SetTarget(ModuleName, GetABI)
	return tx
}

func (tx *ABI) Result() ABIData {
	return tx.Res.(*ABIResult).Data
}

func (tx *ABI) Get(address string) *ABI {
	tx.SetContractAddress(address)
	return tx
}

// Unmarshal converts bytes to a ABIResult.
// Allows ABIResult to implement base.Result
func (r *ABIResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
