package factory

import (
	"Golang-assignment-srikomm/client"
	"Golang-assignment-srikomm/constants"
)

func CryptoClientFactory(cryptoIdentifier string) client.CryptoClientInterface {
	switch cryptoIdentifier {
	case constants.BITCOIN_IDENTIFIER:
		return client.NewCoinDeskClient(constants.COINDESK_ENDPOINT)
	case constants.ETHEREUM_IDENTIFIER:
		return client.NewCryptonatorClient(constants.CRYPTONATOR_ENDPOINT)
	default:
		return client.NewCoinDeskClient(constants.COINDESK_ENDPOINT)
	}
}
