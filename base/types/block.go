package types

type (
	BlockDate struct {
		Number    string `json:"blockNumber"`
		Timestamp string `json:"timeStamp"`
	}

	Block struct {
		BlockDate
		Hash  string `json:"hash"`
		Nonce string `json:"nonce"`
	}
)
