package client

import "Golang-assignment-srikomm/models"

type CryptonatorClientInterface interface {
	GetETHCurrentPrice() (models.Crypto, error)
}
