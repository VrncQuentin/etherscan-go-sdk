package blocks

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	Countdown struct {
		*base.Call
	}

	CountdownResult struct {
		base.CallResult
		Data CountdownData `json:"result"`
	}

	CountdownData struct {
		CurrentBlock      string `json:"CurrentBlock"`
		CountdownBlock    string `json:"CountdownBlock"`
		RemainingBlocks   string `json:"RemainingBlock"`
		EstimateTimeInSec string `json:"EstimateTimeInSec"`
	}
)

func NewCountdown(token string) *Countdown {
	tx := &Countdown{
		base.NewCall(token, new(CountdownResult)),
	}
	tx.SetTarget(ModuleName, GetCountdown)
	return tx
}

func (tx *Countdown) Result() CountdownData {
	return tx.Res.(*CountdownResult).Data
}

func (tx *Countdown) Get(block uint64) *Countdown {
	tx.SetBlockNo(block)
	return tx
}

// Unmarshal converts bytes to a CountdownResult.
// Allows CountdownResult to implement base.Result
func (r *CountdownResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
