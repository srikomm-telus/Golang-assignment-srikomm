package models

type CryptoPriceServiceResponse struct {
	Data        map[string]string
	IsFromCache bool
	CryptoName  string
}

func (cpsr CryptoPriceServiceResponse) GetPriceInCurrency(currencyIdentifier string) string {
	return cpsr.Data[currencyIdentifier]
}
