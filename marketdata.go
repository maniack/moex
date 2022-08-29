package trade

import "fmt"

type MarketData interface {
	Open() (float64, error)
	Low() (float64, error)
	High() (float64, error)
	Last() (float64, error)
	Value() (float64, error)
	Size() (float64, error)
	Rates() (open, low, high, last, value, size float64, err error)
}

type marketdata struct {
	data map[string]interface{}
}

func (m marketdata) Open() (float64, error) {
	retval := 0.0

	switch v := m.data["open"].(type) {
	case float32:
	case float64:
		retval = float64(v)
	default:
		return 0.0, fmt.Errorf("marketdata: wrong type for Open: %T instead of float64", v)
	}

	return retval, nil
}

func (m marketdata) Low() (float64, error) {
	retval := 0.0

	switch v := m.data["low"].(type) {
	case float32:
	case float64:
		retval = float64(v)
	default:
		return 0.0, fmt.Errorf("marketdata: wrong type for Low: %T instead of float64", v)
	}

	return retval, nil
}

func (m marketdata) High() (float64, error) {
	retval := 0.0

	switch v := m.data["high"].(type) {
	case float32:
	case float64:
		retval = float64(v)
	default:
		return 0.0, fmt.Errorf("marketdata: wrong type for High: %T instead of float64", v)
	}

	return retval, nil
}

func (m marketdata) Last() (float64, error) {
	retval := 0.0

	switch v := m.data["last"].(type) {
	case float32:
	case float64:
		retval = float64(v)
	default:
		return 0.0, fmt.Errorf("marketdata: wrong type for Last: %T instead of float64", v)
	}

	return retval, nil
}

func (m marketdata) Value() (float64, error) {
	retval := 0.0

	switch v := m.data["value"].(type) {
	case float32:
	case float64:
		retval = float64(v)
	default:
		return 0.0, fmt.Errorf("marketdata: wrong type for Value: %T instead of float64", v)
	}

	return retval, nil
}

func (m marketdata) Size() (float64, error) {
	value, err := m.Value()
	if err != nil {
		return 0.0, err
	}

	last, err := m.Last()
	if err != nil {
		return 0.0, err
	}

	if last == 0 {
		return 0.0, fmt.Errorf("marketdata: division by zero Last price")
	}

	retval := value / last

	return retval, nil
}

func (m marketdata) Rates() (open, low, high, last, value, size float64, err error) {
	open, err = m.Open()
	if err != nil {
		return open, low, high, last, value, size, err
	}

	low, err = m.Low()
	if err != nil {
		return open, low, high, last, value, size, err
	}

	high, err = m.High()
	if err != nil {
		return open, low, high, last, value, size, err
	}

	last, err = m.Last()
	if err != nil {
		return open, low, high, last, value, size, err
	}

	value, err = m.Value()
	if err != nil {
		return open, low, high, last, value, size, err
	}

	size, err = m.Size()
	if err != nil {
		return open, low, high, last, value, size, err
	}

	return open, low, high, last, value, size, nil
}
