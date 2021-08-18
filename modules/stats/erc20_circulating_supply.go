package stats

import (
	"encoding/json"
	"polygonscan/types/queries"
)

type (
	ERC20CirculatingSupply struct {
		*queries.Call
	}

	ERC20CirculatingSupplyResult struct {
		queries.CallResult
		Supply string `json:"result"`
	}
)

func NewERC20CirculatingSupply(token string) *ERC20CirculatingSupply {
	tx := &ERC20CirculatingSupply{
		queries.NewCall(token, new(ERC20CirculatingSupplyResult)),
	}
	tx.SetTarget(ModuleName, GetTokenCirculatingSupply)
	return tx
}

func (tx *ERC20CirculatingSupply) Result() string {
	return tx.Res.(*ERC20CirculatingSupplyResult).Supply
}

func (tx *ERC20CirculatingSupply) Get(address string) *ERC20CirculatingSupply {
	tx.SetContractAddress(address)
	return tx
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement types.Result
func (r *ERC20CirculatingSupplyResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
