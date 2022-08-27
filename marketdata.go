package trade

import (
	"fmt"
	"strings"
)

type Marketdata interface {
	Rates() (string, float64, float64, float64, float64, float64, float64, float64, float64, error)
	Id() string
	Name() string
}

type marketdata struct {
	s    *security
	data map[string]interface{}
}

func NewMarketdata(s *security, columns []string, data []interface{}) Marketdata {
	m := marketdata{
		s:    s,
		data: make(map[string]interface{}),
	}

	for i, column := range columns {
		m.data[strings.ToLower(column)] = data[i]
	}

	return &m
}

func (m *marketdata) Rates() (id string, open, low, high, last, price, today, value, size float64, err error) {
	switch v := m.data["secid"].(type) {
	case string:
		id = v
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for id: %T insterad of string", v)
	}

	switch v := m.data["marketprice"].(type) {
	case float32:
	case float64:
		price = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for open: %T insterad of float64", v)
	}

	switch v := m.data["marketpricetoday"].(type) {
	case float32:
	case float64:
		today = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for open: %T insterad of float64", v)
	}

	switch v := m.data["open"].(type) {
	case float32:
	case float64:
		open = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for open: %T insterad of float64", v)
	}

	switch v := m.data["low"].(type) {
	case float32:
	case float64:
		low = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for low: %T insterad of float64", v)
	}

	switch v := m.data["high"].(type) {
	case float32:
	case float64:
		high = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for high: %T insterad of float64", v)
	}

	switch v := m.data["last"].(type) {
	case float32:
	case float64:
		last = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for last: %T insterad of float64", v)
	}

	switch v := m.data["value"].(type) {
	case float32:
	case float64:
		value = float64(v)
	default:
		return id, open, low, high, last, price, today, value, size, fmt.Errorf("marketdata: wrong type for last: %T insterad of float64", v)
	}

	size = value / last

	return id, open, low, high, last, price, today, value, size, nil
}

func (m *marketdata) Id() string {
	return m.s.Id()
}

func (m *marketdata) Name() string {
	return m.s.Name()
}
