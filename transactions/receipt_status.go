package transactions

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	ReceiptStatus struct {
		*base.Call
	}

	ReceiptResult struct {
		base.CallResult
		Data ReceiptData `json:"result"`
	}

	ReceiptData struct {
		Status string `json:"status"` // 0 = fail, 1 = pass
	}
)

func NewReceiptStatus(token string) *ReceiptStatus {
	tx := &ReceiptStatus{
		base.NewCall(token, new(ReceiptResult)),
	}
	tx.SetTarget(ModuleName, GetReceiptStatus)
	return tx
}

func (tx *ReceiptStatus) Get(txHash string) *ReceiptStatus {
	tx.SetTxHash(txHash)
	return tx
}

// Unmarshal converts bytes to a ReceiptResult.
// Allows ReceiptResult to implement base.Result
func (r *ReceiptResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
