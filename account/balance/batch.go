package balance

import (
    "encoding/json"
    "fmt"
    "polygonscan/account"
    "polygonscan/base"
)

// Doc: https://polygonscan.com/apis#accounts
// TODO: Find out what `tag` represent in the query & whether it is useful for the wrap

type (
    BatchAccount struct {
        *base.Call
    }

    batchResult struct {
        base.CallResult
        Result []accountBalance `json:"result"`
    }

    accountBalance struct {
        Account string `json:"account"`
        Balance string `json:"balance"`
    }
)

const (
    // Batch requests are limited to 20 accounts, according to polygonscan's API.
    maxAccountPerBatch = 20
)


func NewBatchAccount(token string) *BatchAccount {
    return &BatchAccount{
        base.NewCall(token, new(batchResult)),
    }
}

func (tx *BatchAccount) Result() []accountBalance {
    return tx.Res.(*batchResult).Result
}

// Get gets the balance of multiple addresses in a single call
// Up to 20 addresses per batch.
func (tx *BatchAccount) Get(addresses ...string) (*BatchAccount, error) {
    if len(addresses) > maxAccountPerBatch {
        return nil, fmt.Errorf("too many account requested (%d > %d)", len(addresses), maxAccountPerBatch)
    }

    return tx.get(addresses), nil
}

// MustGet gets the balance of multiple addresses in a single call
// Up to 20 addresses per batch, if given more will ignore all the extra addresses.
func (tx *BatchAccount) MustGet(addresses ...string) *BatchAccount {
    if len(addresses) > maxAccountPerBatch {
        addresses = addresses[:maxAccountPerBatch]
    }

    return tx.get(addresses)
}

func (tx *BatchAccount) get(addresses []string) *BatchAccount {
    tx.SetTarget(account.ModuleName, account.BalanceBatch).
        SetAddresses(addresses)

    return tx
}

// Unmarshal converts bytes to a batchResult.
// Allows batchResult to implement base.Result
func (r *batchResult) Unmarshal(data []byte) error {
    return json.Unmarshal(data, r)
}