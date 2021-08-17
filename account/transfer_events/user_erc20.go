package transfer_events

import (
	"polygonscan/account"
	"polygonscan/base"
)

type UserERC20s struct {
	*base.Call
}

func NewUserERC20s(token string) *UserERC20s {
	tx := &UserERC20s{
		base.NewCall(token, new(ERC20Result)),
	}
	tx.SetTarget(account.ModuleName, account.TransferEventsERC20)
	return tx
}

func (tx *UserERC20s) Result() []ERC20 {
	return tx.Res.(*ERC20Result).Result
}

func (tx *UserERC20s) Get(address string) *UserERC20s {
	tx.SetAddress(address)
	return tx
}

func (tx *UserERC20s) GetBetween(address string, begin, end uint64) *UserERC20s {
	tx.Get(address).SetBlockRange(begin, end)
	return tx
}

func (tx *UserERC20s) PaginatedGet(address string, page, maxEntries uint64) *UserERC20s {
	tx.Get(address).Paginate(page, maxEntries)
	return tx
}

func (tx *UserERC20s) PaginatedGetBetween(
	address string,
	begin, end uint64,
	page, maxEntries uint64,
) *UserERC20s {

	tx.GetBetween(address, begin, end).Paginate(page, maxEntries)
	return tx
}
