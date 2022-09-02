package moex

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Boards - board list
type Boards interface {
	getByName(name string) Board
}

type boards struct {
	mkt  *market
	data []*board
}

func (b boards) getByName(name string) Board {
	for _, brd := range b.data {
		if strings.EqualFold(strings.ToLower(brd.name), strings.ToLower(name)) {
			return brd
		}
	}
	return nil
}

// Board - board (trading mode)
type Board interface {
	Securities() ([]Security, error)
	Security(id string) (Security, error)
}

type board struct {
	mkt   *market
	name  string
	title string
}

// Securities - board securities list
func (b board) Securities() ([]Security, error) {
	res, err := b.mkt.eng.ex.c.Get(fmt.Sprintf(securitiesURI, b.mkt.eng.name, b.mkt.name, b.name))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := new(SecuritiesResponse)

	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	secs := []Security{}
	for i := range r.Securities.Data {
		sdata := make(map[string]interface{})
		for k, c := range r.Securities.Columns {
			sdata[strings.ToLower(c)] = r.Securities.Data[i][k]
		}

		mdata := make(map[string]interface{})
		for k, c := range r.MarketData.Columns {
			mdata[strings.ToLower(c)] = r.MarketData.Data[i][k]
		}

		secid := fmt.Sprintf("%s", sdata["secid"])

		secs = append(secs, &security{strings.ToUpper(secid), securitydata{sdata}, marketdata{mdata}})
	}

	return secs, nil
}

// Security - get security by secid (e.g.: (moex.Board).Security("APPL"))
func (b board) Security(id string) (Security, error) {
	secs, err := b.Securities()
	if err != nil {
		return nil, err
	}

	for _, sec := range secs {
		if strings.EqualFold(strings.ToLower(sec.ID()), strings.ToLower(id)) {
			return sec, nil
		}
	}

	return nil, fmt.Errorf("board: security id %s not foud", id)
}
