package transfer_events

import (
	"polygonscan/account"
	"polygonscan/base"
)

type UserERC721s struct {
	*base.Call
}

func NewUserERC721s(token string) *UserERC721s {
	tx := &UserERC721s{
		base.NewCall(token, new(ERC721Result)),
	}
	tx.SetTarget(account.ModuleName, account.TransferEventsERC721)
	return tx
}

func (tx *UserERC721s) Result() []ERC721 {
	return tx.Res.(*ERC721Result).Result
}

func (tx *UserERC721s) Get(address string) *UserERC721s {
	tx.SetAddress(address)

	return tx
}

func (tx *UserERC721s) GetBetween(address string, begin, end uint64) *UserERC721s {
	tx.Get(address).SetBlockRange(begin, end)
	return tx
}

func (tx *UserERC721s) PaginatedGet(address string, page, maxEntries uint64) *UserERC721s {
	tx.Get(address).Paginate(page, maxEntries)
	return tx
}

func (tx *UserERC721s) PaginatedGetBetween(
	address string,
	begin, end uint64,
	page, maxEntries uint64,
) *UserERC721s {

	tx.GetBetween(address, begin, end).Paginate(page, maxEntries)
	return tx
}
