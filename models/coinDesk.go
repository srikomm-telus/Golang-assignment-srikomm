package models

type CoinBaseResponse struct {
	Time       coinBaseResponseTime             `json:"time"`
	Disclaimer string                           `json:"disclaimer"`
	ChartName  string                           `json:"chartName"`
	Bpi        map[string]cryptoPriceInCurrency `json:"bpi"`
}

type coinBaseResponseTime struct {
	Updated    string `json:"updated"`
	UpdatedISO string `json:"updatedISO"`
	UpdatedUK  string `json:"updatedUK"`
}

type cryptoPriceInCurrency struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rateFloat"`
}

func (cbr CoinBaseResponse) GetPriceInCurrency(currencyIdentifier string) string {
	return cbr.Bpi[currencyIdentifier].Rate
}
