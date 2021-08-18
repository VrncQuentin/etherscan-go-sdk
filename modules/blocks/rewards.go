package blocks

import (
	"encoding/json"
	"polygonscan/types/data"
	"polygonscan/types/queries"
)

type (
	Rewards struct {
		*queries.Call
	}

	RewardsResult struct {
		queries.CallResult
		Data RewardData `json:"result"`
	}

	RewardData struct {
		data.BlockDate

		Miner  string `json:"blockMiner"`
		Reward string `json:"blockReward"`

		Uncles                []string `json:"uncles"` //TODO: Find data representation. string is a placeholder
		UnclesInclusionReward string   `json:"unclesInclusionReward"`
	}
)

func NewRewards(token string) *Rewards {
	tx := &Rewards{
		queries.NewCall(token, new(RewardsResult)),
	}
	tx.SetTarget(ModuleName, GetRewards)
	return tx
}

func (tx *Rewards) Result() RewardData {
	return tx.Res.(*RewardsResult).Data
}

func (tx *Rewards) Get(block uint64) *Rewards {
	tx.SetBlockNo(block)
	return tx
}

// Unmarshal converts bytes to a RewardsResult.
// Allows RewardsResult to implement types.Result
func (r *RewardsResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
