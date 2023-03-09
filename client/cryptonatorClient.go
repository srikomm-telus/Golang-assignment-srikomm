package client

import (
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

var cryptonatorClient *CryptonatorClient

type CryptonatorClient struct {
	endpoint string
}

func NewCryptonatorClient(endpoint string) *CryptonatorClient {
	if cryptonatorClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if cryptonatorClient == nil {
			cryptonatorClient = &CryptonatorClient{
				endpoint: endpoint,
			}
		}
	}
	return cryptonatorClient
}

func (c CryptonatorClient) GetCurrentPrice() (models.Crypto, error) {
	response, err := http.Get(c.endpoint)

	if err != nil {
		logger.Error("Error while fetching crypto price from Cryptonator", zap.String(constants.ERROR_MESSAGE, err.Error()))
		return models.Crypto{}, err
	}
	if response.StatusCode != 200 {
		logger.Error("unable to fetch response from CoinDesk API")
		return models.Crypto{}, err
	}

	var cryptonatorResponse models.CryptonatorResponse

	err = json.NewDecoder(response.Body).Decode(&cryptonatorResponse)

	if err != nil {
		logger.Error("Error while decoding response", zap.String(constants.ERROR_MESSAGE, err.Error()))
		return models.Crypto{}, errors.New("error while decoding response")
	}

	return cryptonatorResponseToCryptoConverter(cryptonatorResponse), nil
}

func cryptonatorResponseToCryptoConverter(cyr models.CryptonatorResponse) models.Crypto {
	price := map[string]string{
		constants.USD_CURRENCY_IDENTIFIER: cyr.Ticker.Price,
	}
	return models.NewCrypto(constants.ETHEREUM_IDENTIFIER, price)
}
