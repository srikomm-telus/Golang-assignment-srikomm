package main

import (
	"Golang-assignment-srikomm/client"
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/factory"
	"Golang-assignment-srikomm/models"
	"Golang-assignment-srikomm/store"
	"context"
)

type CryptoPriceService struct {
	downstreamClient client.CryptoClientInterface
	storageClient    store.CryptoStorageInterface
}

func CryptoPriceServiceForTest(downstreamClient client.CryptoClientInterface, storageClient store.CryptoStorageInterface) *CryptoPriceService {
	return &CryptoPriceService{downstreamClient: downstreamClient, storageClient: storageClient}
}

func NewCryptoPriceService(cryptoName string, ctx context.Context) (CryptoPriceService, error) {
	storageClient, err := factory.CryptoStorageFactory(ctx)
	if err != nil {
		logger.Error(err.Error())
		return CryptoPriceService{}, err
	}
	return CryptoPriceService{
		downstreamClient: factory.CryptoClientFactory(cryptoName),
		storageClient:    storageClient,
	}, nil
}

func (cps CryptoPriceService) GetCryptoPrice(cryptoName string) (models.CryptoPriceServiceResponse, error) {
	storedCryptoPrice, err := cps.cryptoPriceFromCache(cryptoName)
	if err == nil {
		return storedCryptoPrice, err
	}
	cryptoLivePrice, err := cps.cryptoPriceFromDownstream()
	if err != nil {
		return models.CryptoPriceServiceResponse{}, err
	}
	cps.updatePriceInCache(cryptoLivePrice)
	return cryptoLivePrice, nil
}

func (cps CryptoPriceService) cryptoPriceFromDownstream() (models.CryptoPriceServiceResponse, error) {
	cryptoPrice, err := cps.downstreamClient.GetCurrentPrice()
	if err != nil {
		return models.CryptoPriceServiceResponse{}, err
	}
	cryptoPriceConverted := cryptoPriceToAPIResponse(cryptoPrice)
	return cryptoPriceConverted, nil
}

func (cps CryptoPriceService) cryptoPriceFromCache(cryptoName string) (models.CryptoPriceServiceResponse, error) {
	crypto, err := cps.storageClient.GetCryptoPrice(cryptoName)
	if err != nil {
		return models.CryptoPriceServiceResponse{}, err
	}
	return newCryptoPriceServiceResponse(crypto), nil
}

func (cps CryptoPriceService) updatePriceInCache(cpsr models.CryptoPriceServiceResponse) {
	_, err := cps.storageClient.SetCryptoPrice(cryptoFromCryptoPriceServiceResponse(cpsr))
	if err != nil {
		logger.Error(err.Error())
		return
	}
}

func newCryptoPriceServiceResponse(crypto models.Crypto) models.CryptoPriceServiceResponse {
	return models.CryptoPriceServiceResponse{
		Data: map[string]string{
			constants.USD_CURRENCY_IDENTIFIER: crypto.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER),
			constants.EUR_CURRENCY_IDENTIFIER: crypto.GetPriceInCurrency(constants.EUR_CURRENCY_IDENTIFIER),
		},
		IsFromCache: true,
		CryptoName:  crypto.GetCryptoName(),
	}
}

func cryptoFromCryptoPriceServiceResponse(response models.CryptoPriceServiceResponse) models.Crypto {
	return models.NewCrypto(response.CryptoName, response.Data)
}

func cryptoPriceToAPIResponse(c models.Crypto) models.CryptoPriceServiceResponse {
	return models.CryptoPriceServiceResponse{
		Data: map[string]string{
			constants.USD_CURRENCY_IDENTIFIER: c.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER),
			constants.EUR_CURRENCY_IDENTIFIER: c.GetPriceInCurrency(constants.EUR_CURRENCY_IDENTIFIER),
		},
		IsFromCache: false,
		CryptoName:  c.GetCryptoName(),
	}
}
