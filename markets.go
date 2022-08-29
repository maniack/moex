package trade

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Markets interface {
	getByName(name string) Market
}

type markets struct {
	eng  *engine
	data []*market
}

func (m markets) getByName(name string) Market {
	for _, mkt := range m.data {
		if strings.EqualFold(strings.ToLower(mkt.name), strings.ToLower(name)) {
			return mkt
		}
	}
	return nil
}

type Market interface {
	Boards() (Boards, error)
	Board(name string) (Board, error)
}

type market struct {
	eng   *engine
	name  string
	title string
}

func (m *market) Boards() (Boards, error) {
	res, err := m.eng.ex.c.Get(fmt.Sprintf(boardsURI, m.eng.name, m.name))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := new(BoardsResponse)

	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	brds := []*board{}
	for _, data := range r.Boards.Data {
		name := fmt.Sprintf("%s", data[2])
		title := fmt.Sprintf("%s", data[3])
		brds = append(brds, &board{m, name, title})
	}

	return &boards{m, brds}, nil
}

func (m *market) Board(name string) (Board, error) {
	brds, err := m.Boards()
	if err != nil {
		return nil, err
	}

	return brds.getByName(name), nil
}
