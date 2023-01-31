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

type CryptonatorClient struct {
	endpoint string
}

func (c CryptonatorClient) GetCurrentPrice() (Crypto, error) {
	response, err := http.Get(CRYPTONATOR_ENDPOINT)

	if err != nil {
		logger.Error(CRYPTONATOR_FETCH_ERROR_MESSAGE, zap.String(ERROR_MESSAGE, err.Error()))
		return Crypto{}, err
	}

	var cryptonatorResponse CryptonatorResponse

	err = json.NewDecoder(response.Body).Decode(&cryptonatorResponse)

	if err != nil {
		logger.Error(ERROR_WHILE_DECODING_RESPONSE, zap.String(ERROR_MESSAGE, err.Error()))
		return Crypto{}, errors.New(ERROR_WHILE_DECODING_RESPONSE)
	}

	return CryptonatorResponseToCryptoConverter(cryptonatorResponse), nil
}
