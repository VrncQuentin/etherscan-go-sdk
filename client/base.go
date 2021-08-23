package client

import "polygonscan/types/request"

type (

    Client struct {
        Network
        Token string
    }

    Network interface {
        URL() string
    }
)

func NewClient(token string, network Network) *Client {
    return &Client{
        Network: network,
        Token: token,
    }
}

func (c *Client) SwitchNetwork(n Network) {
    c.Network = n
}

func (c *Client) Execute(r request.Request) error {
    r.SetToken(c.Token)
    return r.Execute(c.Network.URL())
}