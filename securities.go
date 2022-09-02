package moex

import (
	"fmt"
	"strings"
)

// Security - trading item
type Security interface {
	ID() string
	Name() string
	IsTrading() bool
	Open() (float64, error)
	Low() (float64, error)
	High() (float64, error)
	Last() (float64, error)
	Value() (float64, error)
	Size() (float64, error)
	Rates() (open, low, high, last, value, size float64, err error)
}

type security struct {
	secid string
	securitydata
	marketdata
}

// ID - return security secid string (e.g.: APPL)
func (s security) ID() string {
	return strings.ToUpper(fmt.Sprintf("%s", s.securitydata.data["secid"]))
}

// Name - return security name string (e.g.: Apple Inc.)
func (s security) Name() string {
	return strings.ToUpper(fmt.Sprintf("%s", s.securitydata.data["latname"]))
}

// IsTrading - return true if security status is Trading ("T")
func (s security) IsTrading() bool {
	return strings.EqualFold(strings.ToUpper(fmt.Sprintf("%s", s.marketdata.data["tradingstatus"])), "T")
}

// Size - security lot size
func (s security) Size() (float64, error) {
	retval := 0.0

	switch v := s.securitydata.data["lotsize"].(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case float32:
	case float64:
		retval = float64(v)
	default:
		return 0.0, fmt.Errorf("marketdata: wrong type for Size: %T instead of float64", v)
	}

	return retval, nil
}

type securitydata struct {
	data map[string]interface{}
}
