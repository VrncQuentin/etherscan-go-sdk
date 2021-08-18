package stats

import (
    "encoding/json"
    "polygonscan/base"
)

type (
    MaticSupply struct {
        *base.Call
    }

    MaticSupplyResult struct {
        base.CallResult
        // Given in Wei, to get value in MATIC divide the returned results by 1000000000000000000
        Supply string `json:"result"`
    }
)

func NewMaticSupply(token string) *MaticSupply {
    tx := &MaticSupply{
        base.NewCall(token, new(MaticSupplyResult)),
    }
    tx.SetTarget(ModuleName, GetTokenSupply)
    return tx
}

func (tx *MaticSupply) Result() string {
    return tx.Res.(*MaticSupplyResult).Supply
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement base.Result
func (r *MaticSupplyResult) Unmarshal(data []byte) error {
    return json.Unmarshal(data, r)
}

