package main

import (
	"Golang-assignment-srikomm/client"
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	"Golang-assignment-srikomm/store"
	"context"
)

type CryptoPriceServiceV2 struct {
	downstreamClient client.CryptoClientInterface
	storageClient    store.CryptoStorageInterface
}

func NewCryptoPriceServiceV2(storageClient store.CryptoStorageInterface) *CryptoPriceServiceV2 {
	return &CryptoPriceServiceV2{
		storageClient: storageClient,
	}
}

func NewCryptoPriceServiceV2ForTest(downstreamClient client.CryptoClientInterface, storageClient store.CryptoStorageInterface) *CryptoPriceServiceV2 {
	return &CryptoPriceServiceV2{
		downstreamClient: downstreamClient,
		storageClient:    storageClient,
	}
}

func (cps *CryptoPriceServiceV2) GetCryptoPrice(cryptoName string, ctx context.Context) (*models.CryptoPriceServiceResponse, error) {
	storedCryptoPrice, err := cps.cryptoPriceFromCache(cryptoName, ctx)
	if err == nil {
		return storedCryptoPrice, err
	} else {
		logger.Info(err.Error())
	}
	if ctx.Value(constants.ENVIRONMENT) != constants.TEST {
		cps.setDownstreamClient(cryptoName)
	}
	cryptoLivePrice, err := cps.cryptoPriceFromDownstream()
	if err != nil {
		return nil, err
	}
	cps.updatePriceInCache(*cryptoLivePrice, ctx)
	return cryptoLivePrice, nil
}

func (cps *CryptoPriceServiceV2) cryptoPriceFromDownstream() (*models.CryptoPriceServiceResponse, error) {
	cryptoPrice, err := cps.downstreamClient.GetCurrentPrice()
	if err != nil {
		return nil, err
	}
	cryptoPriceConverted := cryptoPriceToAPIResponse(cryptoPrice)
	return &cryptoPriceConverted, nil
}

func (cps *CryptoPriceServiceV2) cryptoPriceFromCache(cryptoName string, ctx context.Context) (*models.CryptoPriceServiceResponse, error) {
	crypto, err := cps.storageClient.GetCryptoPrice(cryptoName, ctx)
	if err != nil {
		return nil, err
	}
	return newCryptoPriceServiceResponse(crypto), nil
}

func (cps *CryptoPriceServiceV2) updatePriceInCache(cpsr models.CryptoPriceServiceResponse, ctx context.Context) {
	_, err := cps.storageClient.SetCryptoPrice(cryptoFromCryptoPriceServiceResponse(cpsr), ctx)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}

func (cps *CryptoPriceServiceV2) setDownstreamClient(cryptoName string) {
	switch cryptoName {
	case constants.BITCOIN_IDENTIFIER:
		cps.downstreamClient = client.NewCoinDeskClient(constants.COINDESK_ENDPOINT)
	case constants.ETHEREUM_IDENTIFIER:
		cps.downstreamClient = client.NewCryptonatorClient(constants.CRYPTONATOR_ENDPOINT)
	}
}
