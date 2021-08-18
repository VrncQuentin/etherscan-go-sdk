package account

type Action int

const (
	ModuleName = "account"

	Balance Action = iota
	BalanceBatch
	BalanceToken
	TxList
	TxListInternal
	TransferEventsERC20
	TransferEventsERC721
	MinedBlock
	_max
)

var (
	actions = [_max]string{
		Balance:              "balance",
		BalanceBatch:         "balancemulti",
		BalanceToken:         "tokenbalance",
		TxList:               "txlist",
		TxListInternal:       "txlistinternal",
		TransferEventsERC20:  "tokentx",
		TransferEventsERC721: "tokennfttx",
		MinedBlock:           "getminedblocks",
	}
)

func (a Action) String() string {
	return actions[a]
}
