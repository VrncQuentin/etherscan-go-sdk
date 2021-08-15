package account

import (
    "encoding/json"
    "fmt"
    "polygonscan/base"
)

// Doc: https://polygonscan.com/apis#accounts
// TODO: Find out what `tag` represent in the query & whether it is useful for the wrap

type (
    // BalanceResult represent the JSON for GetMaticBalance
    BalanceResult struct {
        base.ResultImplem
        Balance string `json:"result"`
    }

    // BalanceBatchResult represent the JSON for GetMaticBalanceBatch
    BalanceBatchResult struct {
        base.ResultImplem
        Balances []Balance `json:"result"`
    }

    // Balance is the object used to associate each account to its balance by GetMaticBalanceBatch
    Balance struct {
        Account string `json:"account"`
        Balance string `json:"balance"`
    }
)

// GetMaticBalance for a single address
func GetMaticBalance(token, address string) (*BalanceResult, error) {
    q := base.NewQuery(module, actions[balance], token)
    q.Add("address", address)

    r := new(BalanceResult)
    if err := q.Execute(r); err != nil {
        return nil, err
    }

    return r, nil
}

// GetMaticBalanceBatch for multiple addresses in a single call
// Up to 20 addresses per batch.
func GetMaticBalanceBatch(token string, addresses []string) (*BalanceBatchResult, error) {
    maxAccountPerBatch := 20
    if len(addresses) > maxAccountPerBatch {
        return nil, fmt.Errorf("too many account requested (%d > %d)", len(addresses), maxAccountPerBatch)
    }

    q := base.NewQuery(module, actions[balanceBatch], token)
    for _, addr := range addresses {
        q.Add("address", addr)
    }

    r := new(BalanceBatchResult)
    if err := q.Execute(r); err != nil {
        return nil, err
    }
    return r, nil
}

/*
Implementation of base.BaseResult interface for BalanceResult & BalanceBatch.
Simple call to json.Unmarshal()
*/

func (br *BalanceResult) Unmarshal(body []byte) error {
    return json.Unmarshal(body, br)
}

func (bbr *BalanceBatchResult) Unmarshal(body []byte) error {
    return json.Unmarshal(body, bbr)
}