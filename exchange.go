package trade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Exchange interface {
	Securities() ([]Marketdata, error)
}

type exchange struct {
	client *http.Client
	base   string
	tools  map[string]string
}

func NewExchange() Exchange {
	return &exchange{
		client: http.DefaultClient,
		base:   "https://iss.moex.com",
		tools: map[string]string{
			"securities": "iss/engines/stock/markets/shares/securities",
		},
	}
}

func (e *exchange) Securities() ([]Marketdata, error) {
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

	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	marketdata := []Marketdata{}
	c := r.Marketdata.Columns
	for _, d := range r.Marketdata.Data {
		switch v := d[1].(type) {
		case string:
			if strings.EqualFold(v, "TQBR") {
				m, err := NewMarketdata(c, d)
				if err != nil {
					return nil, err
				}

				marketdata = append(marketdata, m)
			}
		}
	}

	return marketdata, nil
}
