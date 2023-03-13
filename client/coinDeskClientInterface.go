package client

import "Golang-assignment-srikomm/models"

type CoinDeskClientInterface interface {
	GetBTCCurrentPrice() (models.Crypto, error)
}
