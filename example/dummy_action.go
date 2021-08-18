package example

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	SomeCall struct {
		*base.Call
	}

	SomeCallResult struct {
		base.CallResult
		//Data SomeCallData `json:"result"`
	}

	SomeCallData struct {
	}
)

func NewSomeCall(token string) *SomeCall {
	tx := &SomeCall{
		base.NewCall(token, new(SomeCallResult)),
	}
	tx.SetTarget(ModuleName, DoSomeCall)
	return tx
}

func (tx *SomeCall) Result() /**/ {
	//return tx.Res.(*SomeCallResult).
}

// Unmarshal converts bytes to a singleResult.
// Allows singleResult to implement base.Result
func (r *SomeCallResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
