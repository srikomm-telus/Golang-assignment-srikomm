package models

type Currency struct {
	currencyIdentifier string
}

func NewCurrency(currencyIdentifier string) Currency {
	return Currency{
		currencyIdentifier: currencyIdentifier,
	}
}
