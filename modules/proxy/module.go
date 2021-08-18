package proxy

// Doc: https://polygonscan.com/apis#proxy

type Action int

const (
	ModuleName = "proxy"

	GetBlockNumber Action = iota
	GetBlockByNumber
	GetBlockTransactionCountByNumber
	GetTransactionByHash
	GetTransactionByBlockNumberAndIndex
	GetTransactionCount
	GetTransactionReceipt
	SendRawTransaction
	Call
	GetCode
	GetStorageAt // Experimental
	GasPrice
	EstimateGas
	_max
)

var (
	actions = [_max]string{
		GetBlockNumber:                      "eth_blockNumber",
		GetBlockByNumber:                    "eth_getBlockByNumber",
		GetBlockTransactionCountByNumber:    "eth_getBlockTransactionCountByNumber",
		GetTransactionByHash:                "eth_getTransactionByHash",
		GetTransactionByBlockNumberAndIndex: "eth_getTransactionByBlockNumberAndIndex",
		GetTransactionCount:                 "eth_getTransactionCount",
		GetTransactionReceipt:               "eth_getTransactionReceipt",
		SendRawTransaction:                  "eth_sendRawTransaction",
		Call:                                "eth_call",
		GetCode:                             "eth_getCode",
		GetStorageAt:                        "eth_getStorageA",
		GasPrice:                            "eth_gasPrice",
		EstimateGas:                         "eth_estimateGas",
	}
)

func (a Action) String() string {
	return actions[a]
}
