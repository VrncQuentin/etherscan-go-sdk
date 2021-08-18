package transfer_events

import (
	"encoding/json"
	"polygonscan/types/data"
	"polygonscan/types/queries"
)

type (
	// TransferEvent represent a token transfer event
	// It is used by specialized data:
	//   - ERC20
	//   - ERC721
	// Both have an additional similar, yet different, field
	// and are defined below.
	TransferEvent struct {
		data.Tx
		data.Token
	}

	// ERC20 represent a Transfer Event (tx) of ERC20
	ERC20 struct {
        TransferEvent
		Value string `json:"value"`
	}

	// ERC721 represent a Transfer Event (tx) of ERC721
	ERC721 struct {
        TransferEvent
		TokenID string `json:"tokenID"`
	}

	// ERC20Result represent the JSON resulting from GetERC20TransferEvents
	// It is, simply, multiple ERC20 Transfer Events (txs)
	// Implements types.Result interface
	ERC20Result struct {
		queries.CallResult
		Result []ERC20 `json:"result"`
	}

	// ERC721Result represent the JSON resulting from GetERC721TransferEvents.
	// It is, simply, multiple ERC721 Transfer Events (txs)
	// Implements types.Result interface
	ERC721Result struct {
		queries.CallResult
		Result []ERC721 `json:"result"`
	}
)

/*
Implementation of types.Res interface for ERC20Result & ERC721Result.
Simple call to json.Unmarshal()
*/

func (txs *ERC20Result) Unmarshal(data []byte) error {
	return json.Unmarshal(data, txs)
}

func (txs *ERC721Result) Unmarshal(data []byte) error {
	return json.Unmarshal(data, txs)
}
