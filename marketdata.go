package trade

import (
	"fmt"
	"strings"
)

type Marketdata interface {
	Id() string
	Rates() (id string, open, low, high, last, value, size float64, err error)
}

type marketdata struct {
	secid string
	data  map[string]interface{}
}

func NewMarketdata(c []string, d []interface{}) (Marketdata, error) {
	data := make(map[string]interface{})
	for i, s := range c {
		data[strings.ToLower(s)] = d[i]
	}

	switch v := d[0].(type) {
	case string:
		return &marketdata{v, data}, nil
	default:
		return nil, fmt.Errorf("wrong Marketdata format")
	}
}

func (m *marketdata) Id() string {
	return m.secid
}

func (m *marketdata) Rates() (id string, open, low, high, last, value, size float64, err error) {
	switch v := m.data["secid"].(type) {
	case string:
		id = v
	default:
		return id, open, low, high, last, value, size, fmt.Errorf("marketdata: wrong type for secid: %T instead of string", v)
	}

	switch v := m.data["open"].(type) {
	case float32:
	case float64:
		open = float64(v)
	default:
		return id, open, low, high, last, value, size, fmt.Errorf("marketdata: wrong type for open: %T instead of float64", v)
	}

	switch v := m.data["low"].(type) {
	case float32:
	case float64:
		low = float64(v)
	default:
		return id, open, low, high, last, value, size, fmt.Errorf("marketdata: wrong type for low: %T instead of float64", v)
	}

	switch v := m.data["high"].(type) {
	case float32:
	case float64:
		high = float64(v)
	default:
		return id, open, low, high, last, value, size, fmt.Errorf("marketdata: wrong type for high: %T instead of float64", v)
	}

	switch v := m.data["last"].(type) {
	case float32:
	case float64:
		last = float64(v)
	default:
		return id, open, low, high, last, value, size, fmt.Errorf("marketdata: wrong type for last: %T instead of float64", v)
	}

	switch v := m.data["value"].(type) {
	case float32:
	case float64:
		value = float64(v)
	default:
		return id, open, low, high, last, value, size, fmt.Errorf("marketdata: wrong type for value: %T instead of float64", v)
	}

	size = value / last

	return id, open, low, high, last, value, size, nil
}
