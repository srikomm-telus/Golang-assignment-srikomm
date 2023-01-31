package client

import (
	. "Golang-assignment-srikomm/constants"
	. "Golang-assignment-srikomm/models"
	. "Golang-assignment-srikomm/util"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

type CoinDeskClient struct {
	endPoint string
}

var logger, _ = zap.NewProduction()

func (CoinDeskClient) GetCurrentPrice() (Crypto, error) {
	response, err := http.Get(COINDESK_ENDPOINT)

	if err != nil {
		logger.Error(COINDESK_FETCH_ERROR_MESSAGE, zap.String(ERROR_MESSAGE, err.Error()))
		return Crypto{}, err
	}

	var coinBaseResponse CoinBaseResponse

	err = json.NewDecoder(response.Body).Decode(&coinBaseResponse)

	if err != nil {
		logger.Error(ERROR_WHILE_DECODING_RESPONSE, zap.String(ERROR_MESSAGE, err.Error()))
		return Crypto{}, errors.New(ERROR_WHILE_DECODING_RESPONSE)
	}

	return CoinBaseResponseToCryptoConverter(coinBaseResponse), nil
}
