package balance

import (
	"encoding/json"
	"polygonscan/modules/account"
	"polygonscan/types/queries"
)

type (
	ERC20 struct {
		*queries.Call
	}

	ERC20Result struct {
		queries.CallResult
		Balance string `json:"result"`
	}
)

func NewERC20(token string) *ERC20 {
	tx := &ERC20{
		queries.NewCall(token, new(ERC20Result)),
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
// Allows ERC20Result to implement types.Result
func (r *ERC20Result) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
