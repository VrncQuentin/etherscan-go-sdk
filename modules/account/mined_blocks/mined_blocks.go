package mined_blocks

import (
	"encoding/json"
	"polygonscan/modules/account"
	"polygonscan/types/queries"
)

type (
	MinedBlock struct {
		*queries.Call
	}

	Result struct {
		queries.CallResult
		Blocks []Block `json:"result"`
	}

	Block struct {
		Number    string `json:"blockNumber"`
		Timestamp string `json:"timeStamp"`
		Reward    string `json:"blockReward"`
	}
)

func NewMinedBlock(token string) *MinedBlock {
	tx := &MinedBlock{
		queries.NewCall(token, new(Result)),
	}
	tx.SetTarget(account.ModuleName, account.MinedBlock).
		// Only full blocks, no other possible values
		Add("blocktype", "blocks")

	return tx
}

func (tx *MinedBlock) Result() []Block {
	return tx.Res.(*Result).Blocks
}

func (tx *MinedBlock) Get(address string) *MinedBlock {
	tx.SetAddress(address)
	return tx
}

func (tx *MinedBlock) PaginatedGet(address string, page, maxEntries uint64) *MinedBlock {
	tx.Get(address).Paginate(page, maxEntries)
	return tx
}

// Unmarshal converts bytes to a Result.
// Allows Result to implement types.Result
func (r *Result) Unmarshal(body []byte) error {
	return json.Unmarshal(body, r)
}
