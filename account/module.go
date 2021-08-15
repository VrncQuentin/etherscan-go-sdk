package account

type Action int

const (
	ModuleName = "account"

	Balance Action = iota
	BalanceBatch
	TxList
	TxListInternal
	TxListInternalByHash
	TxListInternalByRangeBlock
	TransferEventsERC20
	TransferEventsERC721
	MinedBlock
	_max
)

var (
	actions = [_max]string{
		Balance:                    "Balance",
		BalanceBatch:               "balancemulti",
		TxList:                     "txlist",
		TxListInternal:             "txlistinternal",
		TxListInternalByHash:       "txlistinternal",
		TxListInternalByRangeBlock: "txlistinternal",
		TransferEventsERC20:        "tokentx",
		TransferEventsERC721:       "tokennfttx",
		MinedBlock:                 "getminedblocks",
	}
)

func (a Action) String() string {
	return actions[a]
}