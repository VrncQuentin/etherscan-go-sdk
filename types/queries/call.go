package queries

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Call struct {
	*Query
	Res Result
}

func NewCall(token string, output Result) *Call {
	return &Call{
		Query: NewQuery(token),
		Res:   output,
	}
}

const (
	baseURL = "https://api.polygonscan.com/api"
)

func (c *Call) Execute() error {
	targetURL := baseURL + "?" + c.Encode()
	res, err := http.Get(targetURL)
	if err != nil {
		return fmt.Errorf("GET <%s> failed (%d): %w", targetURL, res.StatusCode, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("<%s> couldn't read body: %w", targetURL, err)
	}

	if err = c.Res.Unmarshal(body); err != nil {
		re := new(resultError)
		if err = re.Unmarshal(body); err == nil {
			return fmt.Errorf("<%s> error: %s", targetURL, re.Result)
		}
		return fmt.Errorf("<%s> couldn't parse body: %w", targetURL, err)
	}
	return nil
}
