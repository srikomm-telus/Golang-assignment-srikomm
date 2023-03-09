package client

import (
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

var lock = &sync.Mutex{}
var coinDeskClient *CoinDeskClient

type CoinDeskClient struct {
	endPoint string
}

func NewCoinDeskClient(endpoint string) *CoinDeskClient {
	if coinDeskClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if coinDeskClient == nil {
			coinDeskClient = &CoinDeskClient{
				endPoint: endpoint,
			}
		}
	}
	return coinDeskClient
}

var logger, _ = zap.NewProduction()

func (c CoinDeskClient) GetCurrentPrice() (models.Crypto, error) {
	response, err := http.Get(c.endPoint)

	if err != nil {
		logger.Error("Error while fetching crypto price from CoinDesk", zap.String(constants.ERROR_MESSAGE, err.Error()))
		return models.Crypto{}, err
	}
	if response.StatusCode != 200 {
		logger.Error("unable to fetch response from CoinDesk API")
		return models.Crypto{}, err
	}

	var coinBaseResponse models.CoinBaseResponse

	err = json.NewDecoder(response.Body).Decode(&coinBaseResponse)

	if err != nil {
		logger.Error("error while decoding response", zap.String(constants.ERROR_MESSAGE, err.Error()))
		return models.Crypto{}, errors.New("error while decoding response")
	}

	return coinBaseResponseToCryptoConverter(coinBaseResponse), nil
}

func (c CoinDeskClient) SetEndpoint(endpoint string) {
	c.endPoint = endpoint
}

func coinBaseResponseToCryptoConverter(cbr models.CoinBaseResponse) models.Crypto {
	price := map[string]string{
		constants.USD_CURRENCY_IDENTIFIER: cbr.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER),
		constants.EUR_CURRENCY_IDENTIFIER: cbr.GetPriceInCurrency(constants.EUR_CURRENCY_IDENTIFIER),
	}
	return models.NewCrypto(constants.BITCOIN_IDENTIFIER, price)
}
