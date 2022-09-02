package moex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Exchange interface {
	Engines() (Engines, error)
	Engine(name string) (Engine, error)
}

type exchange struct {
	c *http.Client
}

func NewExchange() Exchange {
	return &exchange{
		c: http.DefaultClient,
	}
}

func (e *exchange) Engines() (Engines, error) {
	res, err := e.c.Get(enginesURI)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := new(EnginesResponse)

	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	engs := []*engine{}
	for _, data := range r.Engines.Data {
		name := fmt.Sprintf("%s", data[1])
		title := fmt.Sprintf("%s", data[2])
		engs = append(engs, &engine{e, name, title})
	}

	return &engines{e, engs}, nil
}

func (e *exchange) Engine(name string) (Engine, error) {
	engs, err := e.Engines()
	if err != nil {
		return nil, err
	}

	return engs.getByName(name), nil
}
