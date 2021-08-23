package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Implem struct {
	Params url.Values
	Res Result
}

func NewRequest(output Result) *Implem {
	return &Implem{
		Params: make(url.Values),
		Res:   output,
	}
}

func (r *Implem) Execute(targerURL string) error {
	targetURL := targerURL + "?" + r.Params.Encode()
	res, err := http.Get(targetURL)
	if err != nil {
		return fmt.Errorf("GET <%s> failed (%d): %w", targetURL, res.StatusCode, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("<%s> couldn't read body: %w", targetURL, err)
	}

	if err = r.Res.Unmarshal(body); err != nil {
		re := new(resultError)
		if err = re.Unmarshal(body); err == nil {
			return fmt.Errorf("<%s> error: %s", targetURL, re.Result)
		}
		return fmt.Errorf("<%s> couldn't parse body: %w", targetURL, err)
	}
	return nil
}

func (r *Implem) SetToken(token string) *Implem {
	r.Params.Set("apikey", token)
	return r
}

func (r *Implem) SetTarget(module string, action Action) *Implem {
	r.Params.Set("module", module)
	r.Params.Set("action", action.String())
	return r
}


/* Address parameters */

func (r *Implem) SetAddress(address string) *Implem {
	r.Params.Set("address", address)
	return r
}

func (r *Implem) SetAddresses(addresses ...string) *Implem {
	for _, address := range addresses {
		r.Params.Add("address", address)
	}
	return r
}

func (r *Implem) ClearAddresses() *Implem {
	r.Params.Del("address")
	return r
}

func (r *Implem) SetContract(address string) *Implem {
	r.Params.Set("contractaddress", address)
	return r
}


func (r *Implem) SetTxHash(hash string) *Implem {
	r.Params.Set("txhash", hash)
	return r
}

func (r *Implem) SetBlock(block uint) *Implem {
	r.Params.Set("blockno", fmt.Sprintf("%d", block))
	return r
}

func (r *Implem) SetTimestamp(t string) *Implem {
	r.Params.Set("timestamp", t)
	return r
}

var (
	knownTags = []string{
		"earliest",
		"pending",
		"latest",
	}
)

func (r *Implem) SetTag(tag string) *Implem {
	r.Params.Set("tag", tag)
	return r
}

func (r *Implem) Between(begin, end uint64) *Implem {
	r.Params.Set("startblock", fmt.Sprintf("%d", begin))
	r.Params.Set("endblock", fmt.Sprintf("%d", end))
	return r
}

func (r *Implem) Paginate(page, maxEntries uint64) *Implem {
	r.Params.Set("page", fmt.Sprintf("%d", page))
	r.Params.Set("offset", fmt.Sprintf("%d", maxEntries))
	return r
}

func (r *Implem) SortAsc() *Implem {
	return r.sort("asc")
}

func (r *Implem) SortDesc() *Implem {
	return r.sort("desc")
}

func (r *Implem) sort(how string) *Implem {
	r.Params.Set("sort", how)
	return r
}
