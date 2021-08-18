package contracts

type Action int

const (
	ModuleName = "contract"

	GetABI Action = iota
	GetSourceCode

	// TODO: Implem
	AskContractVerification
	CheckVerification
	_max
)

var (
	actions = [_max]string{
		GetABI:                  "getabi",
		GetSourceCode:           "getsourcecode",
		AskContractVerification: "verifysourcecode",
		CheckVerification:       "checkverifystatus",
	}
)

func (a Action) String() string {
	return actions[a]
}
