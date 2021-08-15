package account

import (
	"encoding/json"
	"polygonscan/base"
)

// Doc: https://polygonscan.com/apis#accounts

type (
	// ERC20TransferEventsResult represent the JSON resulting from GetERC20TransferEvents
	// It is, simply, multiple ERC20 Transfer Events (txs)
	// Implements base.Result interface
	ERC20TransferEventsResult struct {
		base.ResultImplem
		Result []ERC20TransferEventResult `json:"result"`
	}

	// ERC721TransferEventsResult represent the JSON resulting from GetERC721TransferEvents.
	// It is, simply, multiple ERC721 Transfer Events (txs)
	// Implements base.Result interface
	ERC721TransferEventsResult struct {
		base.ResultImplem
		Result []ERC721TransferEventResult `json:"result"`
	}

	// ERC20TransferEventResult represent a Transfer Event (tx) of ERC20
	ERC20TransferEventResult struct {
		BaseTransferEventResult
		Value string `json:"value"`
	}

	// ERC721TransferEventResult represent a Transfer Event (tx) of ERC721
	ERC721TransferEventResult struct {
		BaseTransferEventResult
		TokenID string `json:"tokenID"`
	}

	// BaseTransferEventResult represent a Transfer Event (a transaction).
	// TransferEvents between ERC20 & ERC721 are fairly similar hence the usage of a common type.
	BaseTransferEventResult struct {
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
        TokenName string `json:"tokenName"`
        TokenSymbol string `json:"tokenSymbol"`
		TokenDecimal string `json:"tokenDecimal"`

		// Transaction Result data
		Input         string `json:"input"`
		Confirmations string `json:"confirmations"`
	}
)

/* Base Queries */
/* --- User Token Transfers --- */

// GetERC20TransferEvents returns a list of ERC20 Transfer Events (txs) for the given address.
// Address: wallet account
// Maximum: last 10k transactions
func GetERC20TransferEvents(token, address string) (*base.Query, *ERC20TransferEventsResult) {
	q := base.NewQuery(module, actions[transferEventsERC20], token).
		SetAddress(address)
	return q, new(ERC20TransferEventsResult)
}

// GetERC721TransferEvents returns a list of ERC721 Transfer Events (txs) for the given address.
// Address: wallet account
// Maximum: last 10k transactions
func GetERC721TransferEvents(token, address string) (*base.Query, *ERC721TransferEventsResult) {
	q := base.NewQuery(module, actions[transferEventsERC20], token).
		SetAddress(address)
	return q, new(ERC721TransferEventsResult)
}

/* --- Contract Token Transfers --- */

// GetContractERC20TransferEvents returns a list of ERC20 Transfer Events (txs) for the given address.
// Address: contract account
// Maximum: last 10k transactions
func GetContractERC20TransferEvents(token, address string) (*base.Query, *ERC20TransferEventsResult) {
	q := base.NewQuery(module, actions[transferEventsERC20], token).
		SetContractAddress(address)
	return q, new(ERC20TransferEventsResult)
}

// GetContractERC721TransferEvents returns a list of ERC20 Transfer Events (txs) for the given address.
// Address: contract account
// Maximum: last 10k transactions
func GetContractERC721TransferEvents(token, address string) (*base.Query, *ERC721TransferEventsResult) {
	q := base.NewQuery(module, actions[transferEventsERC20], token).
		SetContractAddress(address)
	return q, new(ERC721TransferEventsResult)
}

/* Paginated Queries */
/* --- User Token Transfers --- */

// GetERC20TransferEventsPaginated returns a list of ERC20 Transfer Events (txs) for the given address, with pagination.
// Address: wallet
func GetERC20TransferEventsPaginated(token, address string, page, maxRecords uint64) (*base.Query, *ERC20TransferEventsResult) {
	q, r := GetERC20TransferEvents(token, address)
	q.Paginate(page, maxRecords)

	return q, r
}

// GetERC721TransferEventsPaginated returns a list of ERC721 Transfer Events (txs) for the given address, with pagination.
// Address: wallet
func GetERC721TransferEventsPaginated(token, address string, page, maxRecords uint64) (*base.Query, *ERC721TransferEventsResult) {
	q, r := GetERC721TransferEvents(token, address)
	q.Paginate(page, maxRecords)

	return q, r
}

/* --- Contract Token Transfers --- */

// GetContractERC20TransferEventsPaginated returns a list of ERC20 Transfer Events (txs) for the given address, with pagination.
// Address: contract
func GetContractERC20TransferEventsPaginated(token, address string, page, maxRecords uint64) (*base.Query, *ERC20TransferEventsResult) {
	q, r := GetContractERC20TransferEvents(token, address)
	q.Paginate(page, maxRecords)

	return q, r
}

// GetContractERC721TransferEventsPaginated returns a list of ERC721 Transfer Events (txs) for the given address, with pagination.
// Address: contract
func GetContractERC721TransferEventsPaginated(token, address string, page, maxRecords uint64) (*base.Query, *ERC721TransferEventsResult) {
	q, r := GetContractERC721TransferEvents(token, address)
	q.Paginate(page, maxRecords)

	return q, r
}


/*
Implementation of base.BaseResult interface for ERC20TransferEventsResult & ERC721TransferEventsResult.
Simple call to json.Unmarshal()
*/

func (r *ERC20TransferEventsResult) Unmarshal(body []byte) error {
	return json.Unmarshal(body, r)
}

func (r *ERC721TransferEventsResult) Unmarshal(body []byte) error {
	return json.Unmarshal(body, r)
}
