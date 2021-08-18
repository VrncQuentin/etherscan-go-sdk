package transactions

import (
	"encoding/json"
	"polygonscan/types/queries"
)

type (
	ReceiptStatus struct {
		*queries.Call
	}

	ReceiptResult struct {
		queries.CallResult
		Data ReceiptData `json:"result"`
	}

	ReceiptData struct {
		Status string `json:"status"` // 0 = fail, 1 = pass
	}
)

func NewReceiptStatus(token string) *ReceiptStatus {
	tx := &ReceiptStatus{
		queries.NewCall(token, new(ReceiptResult)),
	}
	tx.SetTarget(ModuleName, GetReceiptStatus)
	return tx
}

func (tx *ReceiptStatus) Get(txHash string) *ReceiptStatus {
	tx.SetTxHash(txHash)
	return tx
}

// Unmarshal converts bytes to a ReceiptResult.
// Allows ReceiptResult to implement types.Result
func (r *ReceiptResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
