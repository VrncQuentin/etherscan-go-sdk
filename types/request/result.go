package request

import "encoding/json"

type (

	Status struct {
		Code    string `json:"status"`
		Message string `json:"message"`
	}

	resultError struct {
		Status
		Result string `json:"result"`
	}
)

func (r *resultError) Unmarshal(body []byte) error {
	return json.Unmarshal(body, r)
}
