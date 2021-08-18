package stats

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	MaticPrice struct {
		*base.Call
	}

	MaticPriceResult struct {
		base.CallResult
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
		base.NewCall(token, new(MaticPriceResult)),
	}
	tx.SetTarget(ModuleName, GetTokenSupply)
	return tx
}

func (tx *MaticPrice) Result() PriceData {
	return tx.Res.(*MaticPriceResult).Data
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement base.Result
func (r *MaticPriceResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
