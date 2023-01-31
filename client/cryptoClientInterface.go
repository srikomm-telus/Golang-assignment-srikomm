package client

import "Golang-assignment-srikomm/models"

type CryptoClientInterface interface {
	GetCurrentPrice() (models.Crypto, error)
}
