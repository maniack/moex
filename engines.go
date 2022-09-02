package moex

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Engines interface {
	getByName(name string) Engine
}

type engines struct {
	exc   *exchange
	data []*engine
}

func (e engines) getByName(name string) Engine {
	for _, eng := range e.data {
		if strings.EqualFold(strings.ToLower(eng.name), strings.ToLower(name)) {
			return eng
		}
	}
	return nil
}

type Engine interface {
	Markets() (Markets, error)
	Market(name string) (Market, error)
}

type engine struct {
	ex    *exchange
	name  string
	title string
}

func (e *engine) Markets() (Markets, error) {
	res, err := e.ex.c.Get(fmt.Sprintf(marketsURI, e.name))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := new(MarketsResponse)

	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	mkts := []*market{}
	for _, data := range r.Markets.Data {
		name := fmt.Sprintf("%s", data[1])
		title := fmt.Sprintf("%s", data[2])
		mkts = append(mkts, &market{e, name, title})
	}

	return &markets{e, mkts}, nil
}

func (e *engine) Market(name string) (Market, error) {
	mkts, err := e.Markets()
	if err != nil {
		return nil, err
	}

	return mkts.getByName(name), nil
}
