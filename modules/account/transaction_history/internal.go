package transaction_history

import (
	"encoding/json"
	"polygonscan/modules/account"
	"polygonscan/types/data"
	"polygonscan/types/queries"
)

type (
	Internal struct {
		*queries.Call
	}

	InternalResult struct {
		queries.CallResult
		Txs []InternalTx `json:"result"`
	}

	InternalTx struct {
		data.Tx

		Value   string `json:"value"`
		Input   string `json:"input"`
		Type    string `json:"type"`
		TraceID string `json:"traceId"`
		IsError string `json:"isError"`
		ErrCode string `json:"errCode"`
	}
)

func NewInternal(token string) *Internal {
	tx := &Internal{
		queries.NewCall(token, new(InternalResult)),
	}
	tx.SetTarget(account.ModuleName, account.TxListInternal)
	return tx
}

func (tx *Internal) Result() []InternalTx {
	return tx.Res.(*InternalResult).Txs
}

// GetByTxHash returns an array with 1 transaction, corresponding to the given hash
func (tx *Internal) GetByTxHash(txHash string) *Internal {
	tx.SetTxHash(txHash)
	return tx
}

// GetByAddress returns a list of transactions, corresponding to the given address,
// up to 10,000 entries
func (tx *Internal) GetByAddress(address string) *Internal {
	tx.SetAddress(address)
	return tx
}

func (tx *Internal) PaginatedGetByAddress(address string, page, maxEntries uint64) *Internal {
	tx.GetByAddress(address).Paginate(page, maxEntries)
	return tx
}

// GetBetween return a list of transactions between the given block,
// up to 10,000 entries
func (tx *Internal) GetBetween(begin, end uint64) *Internal {
	tx.SetBlockRange(begin, end)
	return tx
}

func (tx *Internal) PaginatedGetBetween(
	begin, end uint64,
	page, maxEntries uint64,
) *Internal {

	tx.GetBetween(begin, end).Paginate(page, maxEntries)
	return tx
}

// Unmarshal converts bytes to a InternalResult.
// Allows InternalResult to implement types.Result
func (r *InternalResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
