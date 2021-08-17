package balance

import (
    "encoding/json"
    "polygonscan/account"
    "polygonscan/base"
)

type (
    ERC20 struct {
        *base.Call
    }

    ERC20Result struct {
        base.CallResult
        Balance string `json:"result"`
    }
)

func NewERC20(token string) *ERC20 {
    tx := &ERC20{
        base.NewCall(token, new(ERC20Result)),
    }
    tx.SetTarget(account.ModuleName, account.BalanceToken)
    return tx
}

func (tx *ERC20) Result() string {
    return tx.Res.(*ERC20Result).Balance
}

func (tx *ERC20) Get(account, token string) *ERC20 {
    tx.SetAddress(account).SetContractAddress(token)
    return tx
}

// Unmarshal converts bytes to a ERC20Result.
// Allows ERC20Result to implement base.Result
func (r *ERC20Result) Unmarshal(data []byte) error {
    return json.Unmarshal(data, r)
}