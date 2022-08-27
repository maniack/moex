package trade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Exchange interface {
	Securities() ([]Security, error)
}

type exchange struct {
	client *http.Client
	base   string
	tools  map[string]string
}

func NewExchange() Exchange {
	return &exchange{
		client: &http.Client{},
		base:   "https://iss.moex.com",
		tools: map[string]string{
			"securities": "iss/engines/stock/markets/shares/securities",
		},
	}
}

func (e *exchange) Securities() ([]Security, error) {
	url, err := url.Parse(fmt.Sprintf("%s/%s.json", e.base, e.tools["securities"]))
	if err != nil {
		return nil, err
	}

	r := new(response)

	res, err := e.client.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	securities := []Security{}
	for _, s := range r.Securities.Data {
		securities = append(securities, NewSecurity(e, r.Securities.Columns, s))
	}

	return securities, nil
}
