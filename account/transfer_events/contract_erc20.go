package transfer_events

import (
	"polygonscan/account"
	"polygonscan/base"
)

type ContractERC20s struct {
	*base.Call
}

func NewContractERC20s(token string) *ContractERC20s {
	tx := &ContractERC20s{
		base.NewCall(token, new(ERC20Result)),
	}
	tx.SetTarget(account.ModuleName, account.TransferEventsERC20)
	return tx
}

func (tx *ContractERC20s) Result() []ERC20 {
	return tx.Res.(*ERC20Result).Result
}

func (tx *ContractERC20s) Get(address string) *ContractERC20s {
	tx.SetContractAddress(address)
	return tx
}

func (tx *ContractERC20s) GetBetween(address string, begin, end uint64) *ContractERC20s {
	tx.Get(address).SetBlockRange(begin, end)
	return tx
}

func (tx *ContractERC20s) PaginatedGet(address string, page, maxEntries uint64) *ContractERC20s {
	tx.Get(address).Paginate(page, maxEntries)
	return tx
}

func (tx *ContractERC20s) PaginatedGetBetween(
	address string,
	begin, end uint64,
	page, maxEntries uint64,
) *ContractERC20s {

	tx.GetBetween(address, begin, end).Paginate(page, maxEntries)
	return tx
}
