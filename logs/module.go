package logs

// Doc: https://polygonscan.com/apis#logs

type Action int

const (
	ModuleName = "logs"

	GetLogs Action = iota
	_max
)

var (
	actions = [_max]string{
		GetLogs: "getLogs",
	}
)

func (a Action) String() string {
	return actions[a]
}
