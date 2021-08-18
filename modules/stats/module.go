package stats

type Action int

const (
	ModuleName = "stats"

	GetTokenSupply Action = iota
	GetTokenCirculatingSupply
	GetMaticSupply
	GetMaticPrice
	_max
)

var (
	actions = [_max]string{
		GetTokenSupply:            "tokensupply",
		GetTokenCirculatingSupply: "tokenCsupply",
		GetMaticSupply:            "maticsupply",
		GetMaticPrice:             "maticprice",
	}
)

func (a Action) String() string {
	return actions[a]
}
