package etherscan

import "polygonscan/client"

type (
    Client = client.Client
    NetworkID uint
)

const (
    Mainnet NetworkID = iota
    Goerli
    Kovan
    Rinkeby
    Ropsten
    _max
)

// NewClient creates a client for Etherscan to query Ethereum's mainnet
func NewClient(token string) *Client {
    return client.NewClient(token, Mainnet)
}

/*
The following functions create clients for Etherscan to query Ethereum's testnets
*/

func NewGoerliClient(token string) *Client {
    return client.NewClient(token, Goerli)
}

func NewKovanClient(token string) *Client {
    return client.NewClient(token, Kovan)
}

func NewRinkebyClient(token string) *Client {
    return client.NewClient(token, Rinkeby)
}

func NewRopstenClient(token string) *Client {
    return client.NewClient(token, Ropsten)
}


// URL implements the client.Network interface for NetworkID
// This ensures only valid URL will be used
func (n NetworkID) URL() string {
    networks := [_max]string{
        "https://api.etherscan.io/api",
        "https://api-goerli.etherscan.io/api",
        "https://api-kovan.etherscan.io/api",
        "https://api-rinkeby.etherscan.io/api",
        "https://api-ropsten.etherscan.io/api",
    }
    if n >= _max {
        return networks[Mainnet]
    }
    return networks[n]
}