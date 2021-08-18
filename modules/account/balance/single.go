package balance

import (
	"encoding/json"
	"polygonscan/modules/account"
	"polygonscan/types/queries"
)

// Doc: https://polygonscan.com/apis#accounts
// TODO: Find out what `tag` represent in the query & whether it is useful for the wrap

type (
	SingleAccount struct {
		*queries.Call
	}

	singleResult struct {
		queries.CallResult
		Balance string `json:"result"`
	}
)

func NewSingleAccount(token string) *SingleAccount {
	tx := &SingleAccount{
		queries.NewCall(token, new(singleResult)),
	}
	tx.SetTarget(account.ModuleName, account.Balance)
	return tx
}

func (tx *SingleAccount) Result() string {
	return tx.Res.(*singleResult).Balance
}

func (tx *SingleAccount) Get(address string) *SingleAccount {
	tx.SetAddress(address)
	return tx
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement types.Result
func (r *singleResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
