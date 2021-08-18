package example

type Action int

const (
    ModuleName = "example"

    DoSomeCall Action = iota
    _max
)

var (
    actions = [_max]string{
        DoSomeCall: "",
    }
)

func (a Action) String() string {
    return actions[a]
}
