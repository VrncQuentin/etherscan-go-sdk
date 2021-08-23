package balance

import (
	"encoding/json"
	"errors"
	"polygonscan/modules/account"
	"polygonscan/types/data"
	"polygonscan/types/request"
)

type (
	/* --- End User Types --- */

	// ERC20 is a request specialized for this one:
	// https://docs.etherscan.io/api-endpoints/tokens#get-erc-20-token-account-balance-for-tokencontractaddress
	// It restricts the methods available to only the necessary ones
	ERC20 request.Implem

	// ResultERC20 represents the query's response with proper types
	ResultERC20 struct {
		request.Status
		Balance data.Wei // (10^-16 ETH)
	}

	/* --- Intermediate types --- */

	// responseERC20 represents the query's raw response (everything is a string)
	responseERC20 struct {
		request.Status
		Balance string `json:"result"`
	}
)

func NewERC20(accountAddr, contractAddr string) *ERC20 {
	return (*ERC20)(
		request.NewRequest(new(ResultERC20)).
			SetTarget(account.ModuleName, account.BalanceToken).
			SetAddress(accountAddr).
			SetContract(contractAddr),
	)
}

// Result is a convenient wrapper to retrieve the result of
// in the proper type (and without the request's status)
func (e *ERC20) Result() data.Wei {
	return e.Res.(*ResultERC20).Balance
}

/* --- Allowed parameters --- */

// SetTag accepted value: "earliest", "pending", "latest"
func (e *ERC20) SetTag(tag string) *ERC20 {
	(*request.Implem)(e).SetTag(tag)
	return e
}

// ChangeAddress changes the query parameters to get another account's balance
func (e *ERC20) ChangeAddress(address string) *ERC20 {
	(*request.Implem)(e).SetAddress(address)
	return e
}

// ChangeContract changes the query parameters to get another ERC20's balance
func (e *ERC20) ChangeContract(address string) *ERC20 {
	(*request.Implem)(e).SetContract(address)
	return e
}

/* --- Mandatory methods --- */

// Execute is required to implement the request.Request interface
func (e *ERC20) Execute(targetURL string) error {
	return (*request.Implem)(e).Execute(targetURL)
}

// SetToken is required to implement the request.Request interface
func (e *ERC20) SetToken(token string) {
	(*request.Implem)(e).SetToken(token)
}

// Unmarshal converts bytes to a ResultERC20.
// Allows ResultERC20 to implement request.Result
func (r *ResultERC20) Unmarshal(data []byte) error {
	tmp := responseERC20{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	r.Status = tmp.Status
	_, ok := r.Balance.SetString(tmp.Balance, 10)
	if !ok {
		return errors.New("balance: couldn't convert string to big.Int")
	}
	return nil
}