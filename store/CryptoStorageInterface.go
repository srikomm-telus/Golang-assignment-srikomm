package store

import (
	"Golang-assignment-srikomm/models"
	"context"
)

type CryptoStorageInterface interface {
	SetCryptoPrice(ctx context.Context, crypto models.Crypto) error
	GetCryptoPrice(ctx context.Context, cryptoIdentifier string) (models.Crypto, error)
}
