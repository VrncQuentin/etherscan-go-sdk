package contracts

import (
	"encoding/json"
	"polygonscan/base"
)

type (
	SourceCode struct {
		*base.Call
	}

	SourceCodeResult struct {
		base.CallResult
		Data []SourceCodeData `json:"result"`
	}

	SourceCodeData struct {
		SourceCode           string `json:"SourceCode"`
		ABI                  string `json:"ABI"`
		ContractName         string `json:"ContractName"`
		CompilerVersion      string `json:"CompilerVersion"`
		OptimizationUsed     string `json:"OptimizationUsed"`
		Runs                 string `json:"Runs"`
		ConstructorArguments string `json:"ConstructorArguments"`
		EVMVersion           string `json:"EVMVersion"`
		Library              string `json:"Library"`
		LicenseType          string `json:"LicenseType"`
		Proxy                string `json:"Proxy"`
		Implementation       string `json:"Implementation"`
		SwarmSource          string `json:"SwarmSource"`
	}
)

func NewSourceCode(token string) *SourceCode {
	tx := &SourceCode{
		base.NewCall(token, new(SourceCodeResult)),
	}
	tx.SetTarget(ModuleName, GetSourceCode)
	return tx
}

func (tx *SourceCode) Result() []SourceCodeData {
	return tx.Res.(*SourceCodeResult).Data
}

func (tx *SourceCode) Get(address string) *SourceCode {
	tx.SetContractAddress(address)
	return tx
}

// Unmarshal converts bytes to a SourceCodeResult.
// Allows SourceCodeResult to implement base.Result
func (r *SourceCodeResult) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
