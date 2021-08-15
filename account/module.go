package account

import (
	"polygonscan/base"
)

const (
	module = "account"

	balance base.Action = iota
	balanceBatch
	txList
	txListInternal
	txListInternalByHash
	txListInternalByRangeBlock
	transferEventsERC20
	transferEventsERC721
	minedBlock
	_max

)

var (
	actions = [_max]string{
		balance: "balance",
		balanceBatch: "balancemulti",
		txList: "txlist",
		txListInternal: "txlistinternal",
		txListInternalByHash: "txlistinternal",
		txListInternalByRangeBlock: "txlistinternal",
		transferEventsERC20: "tokentx",
		transferEventsERC721: "tokennfttx",
		minedBlock: "getminedblocks",
	}
)