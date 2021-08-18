package stats

import (
	"encoding/json"
	"polygonscan/types/queries"
)

type (
	ERC20TotalSupply struct {
		*queries.Call
	}

	ERC20TotalSupplyResult struct {
		queries.CallResult
		Supply string `json:"result"`
	}
)

func NewERC20TotalSupply(token string) *ERC20TotalSupply {
	tx := &ERC20TotalSupply{
		queries.NewCall(token, new(ERC20TotalSupplyResult)),
	}
	tx.SetTarget(ModuleName, GetTokenSupply)
	return tx
}

func (tx *ERC20TotalSupply) Result() string {
	return tx.Res.(*ERC20TotalSupplyResult).Supply
}

func (tx *ERC20TotalSupply) Get(address string) *ERC20TotalSupply {
	tx.SetContractAddress(address)
	return tx
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement types.Result
func (r *ERC20TotalSupplyResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
