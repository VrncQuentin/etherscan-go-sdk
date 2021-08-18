package blocks

import (
	"encoding/json"
	"polygonscan/types/queries"
)

type (
	ByTimestamp struct {
		*queries.Call
	}

	ByTimestampResult struct {
		queries.CallResult
		Block string `json:"result"`
	}
)

func NewByTimestamp(token string) *ByTimestamp {
	tx := &ByTimestamp{
		queries.NewCall(token, new(ByTimestampResult)),
	}
	tx.SetTarget(ModuleName, GetByTimestamp)
	return tx
}

func (tx *ByTimestamp) Result() string {
	return tx.Res.(*ByTimestampResult).Block
}

// GetBefore ...
// timestamp must be Unix TS in sec
func (tx *ByTimestamp) GetBefore(timestamp string) *ByTimestamp {
	tx.SetTimestamp(timestamp).
		Add("closest", "before")
	return tx
}

// GetAfter ...
// timestamp must be Unix TS in sec
func (tx *ByTimestamp) GetAfter(timestamp string) *ByTimestamp {
	tx.SetTimestamp(timestamp).
		Add("closest", "after")
	return tx
}

// Unmarshal converts bytes to a ByTimestampResult.
// Allows ByTimestampResult to implement types.Result
func (r *ByTimestampResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
