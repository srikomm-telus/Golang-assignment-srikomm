package models

import . "Golang-assignment-srikomm/constants"

type CryptoPriceServiceResponse struct {
	Data        map[string]string
	IsFromCache bool
	CryptoName  string
}

func NewCryptoPriceServiceResponse(usdPrice string, eurPrice string, isFromCache bool, cryptoName string) CryptoPriceServiceResponse {
	return CryptoPriceServiceResponse{
		Data: map[string]string{
			// TODO extend this by changing this to a for loop to add all the currencies
			USD_CURRENCY_IDENTIFIER: usdPrice,
			EUR_CURRENCY_IDENTIFIER: eurPrice,
		},
		IsFromCache: isFromCache,
		CryptoName:  cryptoName,
	}
}

func (cpsr CryptoPriceServiceResponse) GetPriceInCurrency(currencyIdentifier string) string {
	return cpsr.Data[currencyIdentifier]
}