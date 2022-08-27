package trade

type response struct {
	Securities       data `json:"securities"`
	Marketdata       data `json:"marketdata"`
	Dataversion      data `json:"dataversion"`
	MarketdataYields data `json:"marketdata_yields"`
}

type data struct {
	Metadata map[string]interface{} `json:"metadata"`
	Columns  []string               `json:"columns"`
	Data     [][]interface{}        `json:"data"`
}

type Trader interface{}

type trader struct{}

func NewTrader() Trader {
	return &trader{}
}
