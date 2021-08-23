package polygonscan

import (
    "polygonscan/client"
)

type (
	Client = client.Client
    NetworkID uint
)

const (
    Mainnet NetworkID = iota
    _max
)

// NewClient creates a client for Etherscan to query Ethereum's mainnet
func NewClient(token string) *Client {
    return client.NewClient(token, Mainnet)
}

// URL implements the client.Network interface for NetworkID
// This ensures only valid URL will be used
func (n NetworkID) URL() string {
    networks := [_max]string{
        "https://api.polygonscan.com/api",
    }
    if n >= _max {
        return networks[Mainnet]
    }
    return networks[n]
}
