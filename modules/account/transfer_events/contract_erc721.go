package transfer_events

import (
	"polygonscan/modules/account"
	"polygonscan/types/queries"
)

type ContractERC721s struct {
	*queries.Call
}

func NewContractERC721s(token string) *ContractERC721s {
	tx := &ContractERC721s{
		queries.NewCall(token, new(ERC721Result)),
	}
	tx.SetTarget(account.ModuleName, account.TransferEventsERC721)
	return tx
}

func (tx *ContractERC721s) Result() []ERC721 {
	return tx.Res.(*ERC721Result).Result
}

func (tx *ContractERC721s) Get(address string) *ContractERC721s {
	tx.SetAddress(address)
	return tx
}

func (tx *ContractERC721s) GetBetween(address string, begin, end uint64) *ContractERC721s {
	tx.Get(address).SetBlockRange(begin, end)
	return tx
}

func (tx *ContractERC721s) PaginatedGet(address string, page, maxEntries uint64) *ContractERC721s {
	tx.Get(address).Paginate(page, maxEntries)
	return tx
}

func (tx *ContractERC721s) PaginatedGetBetween(
	address string,
	begin, end uint64,
	page, maxEntries uint64,
) *ContractERC721s {

	tx.GetBetween(address, begin, end).Paginate(page, maxEntries)
	return tx
}
