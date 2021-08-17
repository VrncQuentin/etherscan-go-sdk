package transactions

type Action int

const (
	ModuleName = "transaction"

	GetReceiptStatus Action = iota
	_max
)

var (
	actions = [_max]string{
		GetReceiptStatus: "gettxreceiptstatus",
	}
)

func (a Action) String() string {
	return actions[a]
}
