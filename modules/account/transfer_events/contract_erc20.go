package transfer_events

import (
	"polygonscan/modules/account"
	"polygonscan/types/queries"
)

type ContractERC20s struct {
	*queries.Call
}

func NewContractERC20s(token string) *ContractERC20s {
	tx := &ContractERC20s{
		queries.NewCall(token, new(ERC20Result)),
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
