package models

type Crypto struct {
	cryptoName string
	price      map[string]string
}

func NewCrypto(cryptoName string, price map[string]string) Crypto {
	return Crypto{
		cryptoName: cryptoName,
		price:      price,
	}
}

func (c Crypto) GetPriceInCurrency(currencyIdentifier string) string {
	return c.price[currencyIdentifier]
}

func (c Crypto) GetCryptoName() string {
	return c.cryptoName
}
