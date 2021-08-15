package base

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type (
	Action int
	Module string

	Query struct {
		url.Values
	}
)

const (
	baseURL = "https://api.polygonscan.com/api"
)

func NewQuery(module, action, token string) Query {
	return Query{
		url.Values{
			"module": []string{module},
			"action": []string{action},
			"apikey": []string{token},
		},
	}
}

func (q Query) Paginate(page, maxRecords uint64) Query {
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("offset", fmt.Sprintf("%d", maxRecords))
	return q
}

func (q Query) SetAddress(address string) Query {
	q.Add("address", address)
	return q
}

func (q Query) Execute(r Result) error {
	targetURL := baseURL + "?" + q.Encode()
	res, err := http.Get(targetURL)
	if err != nil {
		return fmt.Errorf("GET <%s> failed (%d): %w", targetURL, res.StatusCode, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("<%s> couldn't read body: %w", targetURL, err)
	}

	if err = r.Unmarshal(body); err != nil {
		re := new(resultError)
		fmt.Println(string(body), err)
		if err = re.Unmarshal(body); err == nil {
			return fmt.Errorf("<%s> error: %s", targetURL, re.Result)
		}
		return fmt.Errorf("<%s> couldn't parse body: %w", targetURL, err)
	}
	return nil
}