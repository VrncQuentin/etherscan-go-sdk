package stats

import (
	"encoding/json"
	"polygonscan/types/queries"
)

type (
	MaticSupply struct {
		*queries.Call
	}

	MaticSupplyResult struct {
		queries.CallResult
		// Given in Wei, to get value in MATIC divide the returned results by 1000000000000000000
		Supply string `json:"result"`
	}
)

func NewMaticSupply(token string) *MaticSupply {
	tx := &MaticSupply{
		queries.NewCall(token, new(MaticSupplyResult)),
	}
	tx.SetTarget(ModuleName, GetTokenSupply)
	return tx
}

func (tx *MaticSupply) Result() string {
	return tx.Res.(*MaticSupplyResult).Supply
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement types.Result
func (r *MaticSupplyResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
