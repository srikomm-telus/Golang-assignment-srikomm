package factory

import (
	"Golang-assignment-srikomm/client"
	"Golang-assignment-srikomm/constants"
)

func CryptoClientFactory(cryptoIdentifier string) client.CryptoClientInterface {
	switch cryptoIdentifier {
	case constants.BITCOIN_IDENTIFIER:
		return client.CoinDeskClient{}
	case constants.ETHEREUM_IDENTIFIER:
		return client.CryptonatorClient{}
	default:
		return client.CoinDeskClient{}
	}
}
