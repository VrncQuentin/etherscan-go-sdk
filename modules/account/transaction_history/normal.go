package transaction_history

import (
	"encoding/json"
	"polygonscan/modules/account"
	"polygonscan/types/data"
	"polygonscan/types/queries"
)

type (
	Normal struct {
		*queries.Call
	}

	NormalResult struct {
		queries.CallResult
		Txs []NormalTx `json:"result"`
	}

	NormalTx struct {
		data.Tx

		BlockHash     string `json:"blockHash"`
		Value         string `json:"value"`
		IsError       string `json:"isError"` // values: 0=No Error, 1=Got Error
		ReceiptStatus string `json:"txreceipt_status"`
		Input         string `json:"input"`
	}
)

func NewNormal(token string) *Normal {
	tx := &Normal{
		queries.NewCall(token, new(NormalResult)),
	}
	tx.SetTarget(account.ModuleName, account.TxList)
	return tx
}

func (tx *Normal) Result() []NormalTx {
	return tx.Res.(*NormalResult).Txs
}

func (tx *Normal) Get(address string) *Normal {
	tx.SetAddress(address)
	return tx
}

func (tx *Normal) GetBetween(address string, begin, end uint64) *Normal {
	tx.Get(address).SetBlockRange(begin, end)
	return tx
}

func (tx *Normal) PaginatedGet(address string, page, maxEntries uint64) *Normal {
	tx.Get(address).Paginate(page, maxEntries)
	return tx
}

func (tx *Normal) PaginatedGetBetween(
	address string,
	begin, end uint64,
	page, maxEntries uint64,
) *Normal {

	tx.GetBetween(address, begin, end).Paginate(page, maxEntries)
	return tx
}

// Unmarshal converts bytes to a NormalResult.
// Allows NormalResult to implement types.Result
func (r *NormalResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
