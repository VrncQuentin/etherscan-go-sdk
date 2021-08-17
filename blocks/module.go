package blocks

type Action int

const (
	ModuleName = "block"

	GetRewards Action = iota
	GetCountdown
	GetByTimestamp
	_max
)

var (
	actions = [_max]string{
		GetRewards:     "getblockreward",
		GetCountdown:   "getblockcountdown",
		GetByTimestamp: "getblocknobytime",
	}
)

func (a Action) String() string {
	return actions[a]
}
