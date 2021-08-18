package queries

import "encoding/json"

type (
	Result interface {
		Unmarshal(body []byte) error
	}

	CallResult struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	resultError struct {
		CallResult
		Result string `json:"result"`
	}
)

func (r *resultError) Unmarshal(body []byte) error {
	return json.Unmarshal(body, r)
}
