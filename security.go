package trade

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Security interface {
	Marketdata() ([]Marketdata, error)
	Id() string
	Name() string
}

type security struct {
	e    *exchange
	data map[string]interface{}
}

func NewSecurity(e *exchange, columns []string, data []interface{}) Security {
	s := security{
		e:    e,
		data: make(map[string]interface{}),
	}

	for i, column := range columns {
		s.data[strings.ToLower(column)] = data[i]
	}

	return &s
}

func (s *security) Marketdata() ([]Marketdata, error) {
	url, err := url.Parse(fmt.Sprintf("%s/%s/%s.json", s.e.base, s.e.tools["securities"], s.data["secid"]))
	if err != nil {
		return nil, err
	}

	r := new(response)

	res, err := s.e.client.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	marketdata := []Marketdata{}
	for _, m := range r.Marketdata.Data {
		marketdata = append(marketdata, NewMarketdata(s, r.Marketdata.Columns, m))
	}

	return marketdata, nil
}

func (s *security) Id() string {
	return fmt.Sprint(s.data["secid"])
}

func (s *security) Name() string {
	return fmt.Sprint(s.data["secname"])
}
