package transfer_events

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	// Token represent a token transfer event
    // It is used by specialized types:
	//   - ERC20
	//   - ERC721
    // Both have an additional similar, yet different, field
    // and are defined below.
	Token struct {
		// Block data
		Block     string `json:"blockNumber"`
		Timestamp string `json:"timeStamp"`
		Hash      string `json:"hash"`
		Nonce     string `json:"nonce"`

		// Participants data
		From            string `json:"from"`
		To              string `json:"to"`
		ContractAddress string `json:"contractAddress"`

		// Transaction data
		TxIndex string `json:"transactionIndex"`
		// -> Gas data
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		GasUsed           string `json:"gasUsed"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		// -> Token
		TokenName    string `json:"tokenName"`
		TokenSymbol  string `json:"tokenSymbol"`
		TokenDecimal string `json:"tokenDecimal"`

		// Transaction Res data
		Input         string `json:"input"`
		Confirmations string `json:"confirmations"`
	}

	// ERC20 represent a Transfer Event (tx) of ERC20
	ERC20 struct {
		Token
		Value string `json:"value"`
	}
	// ERC721 represent a Transfer Event (tx) of ERC721
	ERC721 struct {
		Token
		TokenID string `json:"tokenID"`
	}

	// ERC20Result represent the JSON resulting from GetERC20TransferEvents
	// It is, simply, multiple ERC20 Transfer Events (txs)
	// Implements base.Result interface
	ERC20Result struct {
		base.CallResult
		Result []ERC20 `json:"result"`
	}

	// ERC721Result represent the JSON resulting from GetERC721TransferEvents.
	// It is, simply, multiple ERC721 Transfer Events (txs)
	// Implements base.Result interface
	ERC721Result struct {
		base.CallResult
		Result []ERC721 `json:"result"`
	}
)

/*
Implementation of base.Res interface for ERC20Result & ERC721Result.
Simple call to json.Unmarshal()
*/

func (txs *ERC20Result) Unmarshal(data []byte) error {
	return json.Unmarshal(data, txs)
}

func (txs *ERC721Result) Unmarshal(data []byte) error {
	return json.Unmarshal(data, txs)
}