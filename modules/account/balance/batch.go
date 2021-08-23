package balance

import (
	"encoding/json"
	"errors"
	"fmt"
	"polygonscan/modules/account"
	"polygonscan/types/data"
	"polygonscan/types/request"
)

type (
	/* --- End User Types --- */

	// BatchAccount is a request specialized for this one:
	// https://docs.etherscan.io/api-endpoints/accounts#get-ether-balance-for-multiple-addresses-in-a-single-call
	// It restricts the methods available to only the necessary ones
	BatchAccount request.Implem

	// ResultBA represents the query's response with proper types
	ResultBA struct {
		request.Status
		Accounts []AccountBalance `json:"accounts"`
	}

	// AccountBalance represent an element of the response's array with proper types
	AccountBalance struct {
		Account string `json:"account"`
		Balance data.Wei `json:"balance"`
	}

	/* --- Intermediate types --- */

	// responseBA represents the query's raw response (everything is a string)
	responseBA struct {
		request.Status
		Accounts []responseAccountBalance `json:"result"`
	}

	// responseAccountBalance represent an element of the response's array
	responseAccountBalance struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
	}

)

const (
	// MaxAccountPerBatch represent Etherscan's limit per batch request.
	MaxAccountPerBatch = 20
)

// NewBatchAccount creates a query to get multiple's account balances.
// If more than MaxAccountPerBatch addresses are provided, returns an error.
func NewBatchAccount(addresses ...string) (*BatchAccount, error) {
	if len(addresses) > MaxAccountPerBatch {
		return nil, fmt.Errorf("too many account requested (%d > %d)", len(addresses), MaxAccountPerBatch)
	}
	return (*BatchAccount)(
		request.NewRequest(new(ResultBA)).
			SetTarget(account.ModuleName, account.BalanceBatch).
			SetAddresses(addresses...),
	), nil
}

// MustNewBatchAccount creates a query to get multiple's account balances.
// If more than MaxAccountPerBatch addresses are provided, ignore the extra ones.
func MustNewBatchAccount(addresses ...string) *BatchAccount {
	if len(addresses) > MaxAccountPerBatch {
		addresses = addresses[:MaxAccountPerBatch]
	}
	return (*BatchAccount)(
		request.NewRequest(new(ResultBA)).
			SetTarget(account.ModuleName, account.BalanceBatch).
			SetAddresses(addresses...),
	)
}

// Result is a convenient wrapper to retrieve the result of
// in the proper type (and without the request's status)
func (ba *BatchAccount) Result() []AccountBalance {
	return ba.Res.(*ResultBA).Accounts
}

/* --- Allowed Parameters --- */

// SetTag accepted value: "earliest", "pending", "latest"
func (ba *BatchAccount) SetTag(tag string) *BatchAccount {
	(*request.Implem)(ba).SetTag(tag)
	return ba
}

// ChangeAddresses changes the set of addresses to be queried.
// If more than MaxAccountPerBatch addresses are provided, returns an error.
func (ba *BatchAccount) ChangeAddresses(addresses ...string) (*BatchAccount, error) {
	if len(addresses) > MaxAccountPerBatch {
		return nil, fmt.Errorf("too many account requested (%d > %d)", len(addresses), MaxAccountPerBatch)
	}
	(*request.Implem)(ba).ClearAddresses().SetAddresses(addresses...)
	return ba, nil
}

// MustChangeAddresses changes the set of addresses to be queried.
// If more than MaxAccountPerBatch addresses are provided, ignore the extra ones.
func (ba *BatchAccount) MustChangeAddresses(addresses ...string) *BatchAccount {
	if len(addresses) > MaxAccountPerBatch {
		addresses = addresses[:MaxAccountPerBatch]
	}
	(*request.Implem)(ba).ClearAddresses().SetAddresses(addresses...)
	return ba
}

/* --- Mandatory methods --- */

// Execute is required to implement the request.Request interface
func (ba *BatchAccount) Execute(targetURL string) error {
	return (*request.Implem)(ba).Execute(targetURL)
}

// SetToken is required to implement the request.Request interface
func (ba *BatchAccount) SetToken(token string) {
	(*request.Implem)(ba).SetToken(token)
}

// Unmarshal converts bytes to a ResultSA.
// Allows ResultSA to implement request.Result
func (r *ResultBA) Unmarshal(data []byte) error {
	tmp := responseBA{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	r.Status = tmp.Status
	r.Accounts = make([]AccountBalance, len(tmp.Accounts))
	for i, acc := range tmp.Accounts {
		r.Accounts[i].Account = acc.Account
		_, ok := r.Accounts[i].Balance.SetString(acc.Balance, 10)
		if !ok {
			return errors.New("balance: couldn't convert string to big.Int")
		}
	}
	return nil
}