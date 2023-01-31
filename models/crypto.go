package models

type Crypto struct {
	cryptoName string
	price      map[Currency]string
}

func NewCrypto(cryptoName string, price map[Currency]string) Crypto {
	return Crypto{
		cryptoName: cryptoName,
		price:      price,
	}
}

func (c Crypto) GetPriceInCurrency(currencyIdentifier string) string {
	return c.price[NewCurrency(currencyIdentifier)]
}

func (c Crypto) GetCryptoName() string {
	return c.cryptoName
}
