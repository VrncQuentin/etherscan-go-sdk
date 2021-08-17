package types

type (
	Tx struct {
		Block
		Participants
		Gas
		TxInfo
	}

	Participants struct {
		From            string `json:"from"`
		To              string `json:"to"`
		ContractAddress string `json:"contractAddress"`
	}

	Gas struct {
		Amount         string `json:"gas"`
		Price          string `json:"gasPrice"`
		Used           string `json:"gasUsed"`
		CumulativeUsed string `json:"cumulativeGasUsed"`
	}

	TxInfo struct {
		Index         string `json:"transactionIndex"`
		Input         string `json:"input"`
		Confirmations string `json:"confirmations"`
	}
)
