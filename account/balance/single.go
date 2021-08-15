package balance

import (
    "encoding/json"
    "polygonscan/account"
    "polygonscan/base"
)

// Doc: https://polygonscan.com/apis#accounts
// TODO: Find out what `tag` represent in the query & whether it is useful for the wrap

type (
    SingleAccount struct {
        *base.Call
    }

    singleResult struct {
        base.CallResult
        Balance string `json:"result"`
    }
)

func NewSingleAccount(token string) *SingleAccount {
    return &SingleAccount{
        base.NewCall(token, new(singleResult)),
    }
}

func (tx *SingleAccount) Result() string {
    return tx.Res.(*singleResult).Balance
}

func (tx *SingleAccount) Get(address string) *SingleAccount {
    tx.SetTarget(account.ModuleName, account.Balance).
        SetAddress(address)

    return tx
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement base.Result
func (r *singleResult) Unmarshal(data []byte) error {
    return json.Unmarshal(data, r)
}