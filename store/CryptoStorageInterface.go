package store

import (
	"Golang-assignment-srikomm/models"
	"context"
)

type CryptoStorageInterface interface {
	SetCryptoPrice(crypto models.Crypto, ctx context.Context) (bool, error)
	GetCryptoPrice(cryptoIdentifier string, ctx context.Context) (models.Crypto, error)
}
