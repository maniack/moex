package moex

const (
	iisURI        string = `https://iss.moex.com/iss.json`
	enginesURI    string = `https://iss.moex.com/iss/engines.json`
	engineURI     string = `https://iss.moex.com/iss/engines/%s.json`
	marketsURI    string = `https://iss.moex.com/iss/engines/%s/markets.json`
	marketURI     string = `https://iss.moex.com/iss/engines/%s/markets/%s.json`
	boardsURI     string = `https://iss.moex.com/iss/engines/%s/markets/%s/boards.json`
	boardURI      string = `https://iss.moex.com/iss/engines/%s/markets/%s/boards/%s.json`
	securitiesURI string = `https://iss.moex.com/iss/engines/%s/markets/%s/boards/%s/securities.json`
	securityURI   string = `https://iss.moex.com/iss/engines/%s/markets/%s/boards/%s/securities/%s.json`
)

// Response - MOEX API generic response
type Response struct {
	Metadata map[string]interface{} `json:"metadata"`
	Columns  []string               `json:"columns"`
	Data     [][]interface{}        `json:"data"`
}

// Response - MOEX API IIS response
type IisResponse struct {
	Engines             Response `json:"engines"`
	Markets             Response `json:"markets"`
	Boards              Response `json:"boards"`
	BoardGroups         Response `json:"boardgroups"`
	Durations           Response `json:"durations"`
	SecurityTypes       Response `json:"securitytypes"`
	SecurityGroups      Response `json:"securitygroups"`
	SecurityCollections Response `json:"securitycollections"`
}

// Response - MOEX API engine list response
type EnginesResponse struct {
	Engines Response `json:"engines"`
}

// Response - MOEX API engine response
type EngineResponse struct {
	Engine     Response `json:"engine"`
	TimeTable  Response `json:"timetable"`
	DailyTable Response `json:"dailytable"`
}

// Response - MOEX API market list response
type MarketsResponse struct {
	Markets Response `json:"markets"`
}

// Response - MOEX API market response
type MarketResponse struct {
	Boards           Response `json:"boards"`
	BoardGroups      Response `json:"boardgroups"`
	Securities       Response `json:"securities"`
	MarketData       Response `json:"marketdata"`
	Trades           Response `json:"trades"`
	OrderBook        Response `json:"orderbook"`
	History          Response `json:"history"`
	TradesHistory    Response `json:"trades_hist"`
	MarketDataYields Response `json:"marketdata_yields"`
	TradesYields     Response `json:"trades_yields"`
	HistoryYields    Response `json:"history_yields"`
}

// Response - MOEX API board list response
type BoardsResponse struct {
	Boards Response `json:"boards"`
}

// Response - MOEX API board response
type BoardResponse struct {
	Board Response `json:"board"`
}

// Response - MOEX API securitiy list response
type SecuritiesResponse struct {
	Securities       Response `json:"securities"`
	MarketData       Response `json:"marketdata"`
	DataVersion      Response `json:"dataversion"`
	MarketdataYields Response `json:"marketdata_yields"`
}

// Response - MOEX API securitiy response
type SecurityResponse struct {
	Securities       Response `json:"securities"`
	MarketData       Response `json:"marketdata"`
	DataVersion      Response `json:"dataversion"`
	MarketdataYields Response `json:"marketdata_yields"`
}
