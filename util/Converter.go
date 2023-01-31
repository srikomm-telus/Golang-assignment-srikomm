package Util

import (
	. "Golang-assignment-srikomm/constants"
	. "Golang-assignment-srikomm/models"
)

func CoinBaseResponseToCryptoConverter(cbr CoinBaseResponse) Crypto {
	price := map[Currency]string{
		NewCurrency(USD_CURRENCY_IDENTIFIER): cbr.GetPriceInCurrency(USD_CURRENCY_IDENTIFIER),
		NewCurrency(EUR_CURRENCY_IDENTIFIER): cbr.GetPriceInCurrency(EUR_CURRENCY_IDENTIFIER),
	}
	return NewCrypto(BITCOIN_IDENTIFIER, price)
}

func CryptoPriceToAPIResponse(c Crypto) CryptoPriceServiceResponse {
	return CryptoPriceServiceResponse{
		Data: map[string]string{
			// TODO extend this by changing this to a for loop to add all the currencies
			USD_CURRENCY_IDENTIFIER: c.GetPriceInCurrency(USD_CURRENCY_IDENTIFIER),
			EUR_CURRENCY_IDENTIFIER: c.GetPriceInCurrency(EUR_CURRENCY_IDENTIFIER),
		},
		IsFromCache: PRICE_IS_NOT_FROM_CACHE,
		CryptoName:  c.GetCryptoName(),
	}
}

func CryptonatorResponseToCryptoConverter(cyr CryptonatorResponse) Crypto {
	price := map[Currency]string{
		NewCurrency(USD_CURRENCY_IDENTIFIER): cyr.Ticker.Price,
	}
	return NewCrypto(ETHEREUM_IDENTIFIER, price)
}
