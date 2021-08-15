package base

import "encoding/json"

type (
    Result interface {
        Unmarshal(body []byte) error
    }

    ResultImplem struct {
        Status string `json:"status"`
        Message string `json:"message"`
    }

    resultError struct {
        ResultImplem
        Result string `json:"result"`
    }
)

func (r *resultError) Unmarshal(body []byte) error {
    return json.Unmarshal(body, r)
}