package request

type (
    Request interface {
        SetToken(string)
        Execute(targetURL string) error
    }

    // Action represent the 'action' parameters of request
    // This interface forces modules to clearly defines
    // their actions.
    Action interface {
        String() string
    }

    Result interface {
        Unmarshal(body []byte) error
    }

)
