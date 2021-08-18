package stats

import (
	"encoding/json"
	"polygonscan/types/queries"
)

type (
	MaticPrice struct {
		*queries.Call
	}

	MaticPriceResult struct {
		queries.CallResult
		Data PriceData `json:"result"`
	}

	PriceData struct {
		MaticBTC          string `json:"maticbtc"`
		TimestampMaticBTC string `json:"maticbtc_timestamp"`

		MaticUSD          string `json:"maticusd"`
		TimestampMaticUSD string `json:"maticusd_timestamp"`
	}
)

func NewMaticPrice(token string) *MaticPrice {
	tx := &MaticPrice{
		queries.NewCall(token, new(MaticPriceResult)),
	}
	tx.SetTarget(ModuleName, GetTokenSupply)
	return tx
}

func (tx *MaticPrice) Result() PriceData {
	return tx.Res.(*MaticPriceResult).Data
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement types.Result
func (r *MaticPriceResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
