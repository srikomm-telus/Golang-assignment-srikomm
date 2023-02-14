package store

import "Golang-assignment-srikomm/models"

type CryptoStorageInterface interface {
	SetCryptoPrice(crypto models.Crypto) (bool, error)
	GetCryptoPrice(cryptoIdentifier string) (models.Crypto, error)
}
