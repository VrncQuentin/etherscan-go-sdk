package stats

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	ERC20CirculatingSupply struct {
		*base.Call
	}

	ERC20CirculatingSupplyResult struct {
		base.CallResult
		Supply string `json:"result"`
	}
)

func NewERC20CirculatingSupply(token string) *ERC20CirculatingSupply {
	tx := &ERC20CirculatingSupply{
		base.NewCall(token, new(ERC20CirculatingSupplyResult)),
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
// Allows singleResult to implement base.Result
func (r *ERC20CirculatingSupplyResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
