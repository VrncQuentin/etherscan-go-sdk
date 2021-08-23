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

	// SingleAccount is a request specialized for this one:
	// https://docs.etherscan.io/api-endpoints/accounts#get-ether-balance-for-a-single-address
	// It restricts the methods available to only the necessary ones
	SingleAccount request.Implem

	// ResultSA represents the query's response with proper types
	ResultSA struct {
		request.Status
		Balance data.Wei // (10^-16 ETH)
	}

	/* --- Intermediate types --- */

	// responseSA represents the query's raw response (everything is a string)
	responseSA struct {
		request.Status
		Balance string `json:"result"`
	}
)

func NewSingleAccount(address string) *SingleAccount {
	return (*SingleAccount)(
		request.NewRequest(new(ResultSA)).
			SetTarget(account.ModuleName, account.Balance).
			SetAddress(address),
	)
}

// Result is a convenient wrapper to retrieve the result of
// in the proper type (and without the request's status)
func (sa *SingleAccount) Result() data.Wei {
	return sa.Res.(*ResultSA).Balance
}

/* --- Allowed parameters --- */

// SetTag accepted value: "earliest", "pending", "latest"
func (sa *SingleAccount) SetTag(tag string) *SingleAccount {
	(*request.Implem)(sa).SetTag(tag)
	return sa
}

// ChangeAddress changes the query parameters to get another account's balance
func (sa *SingleAccount) ChangeAddress(address string) *SingleAccount {
	(*request.Implem)(sa).SetAddress(address)
	return sa
}

/* --- Mandatory methods --- */

// Execute is required to implement the request.Request interface
func (sa *SingleAccount) Execute(targetURL string) error {
	return (*request.Implem)(sa).Execute(targetURL)
}

// SetToken is required to implement the request.Request interface
func (sa *SingleAccount) SetToken(token string) {
	(*request.Implem)(sa).SetToken(token)
}

// Unmarshal converts bytes to a ResultSA.
// Allows ResultSA to implement request.Result
func (r *ResultSA) Unmarshal(data []byte) error {
	tmp := responseSA{}
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