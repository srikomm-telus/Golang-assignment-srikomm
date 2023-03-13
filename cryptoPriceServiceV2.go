package main

import (
	"Golang-assignment-srikomm/client"
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	"Golang-assignment-srikomm/store"
	"context"
)

type CryptoPriceServiceV2 struct {
	coinDeskClient    client.CoinDeskClientInterface
	cryptonatorClient client.CryptonatorClientInterface
	storageClient     store.CryptoStorageInterface
}

func NewCryptoPriceServiceV2(coinDeskClient client.CoinDeskClientInterface,
	cryptonatorClient client.CryptonatorClientInterface, redisClient *store.RedisClient) *CryptoPriceServiceV2 {
	return &CryptoPriceServiceV2{
		coinDeskClient:    coinDeskClient,
		cryptonatorClient: cryptonatorClient,
		storageClient: store.CryptoCacheStorage{
			CacheClient: redisClient,
		},
	}
}

func NewCryptoPriceServiceV2ForTest(coinDeskClient client.CoinDeskClientInterface,
	cryptonatorClient client.CryptonatorClientInterface, storageClient store.CryptoStorageInterface) *CryptoPriceServiceV2 {
	return &CryptoPriceServiceV2{
		coinDeskClient:    coinDeskClient,
		cryptonatorClient: cryptonatorClient,
		storageClient:     storageClient,
	}
}

func (cps *CryptoPriceServiceV2) GetCryptoPrice(ctx context.Context, cryptoName string) (models.CryptoPriceServiceResponse, error) {
	storedCryptoPrice, err := cps.cryptoPriceFromCache(ctx, cryptoName)
	if err == nil {
		return storedCryptoPrice, err
	} else {
		logger.Info(err.Error())
	}
	cryptoLivePrice, err := cps.cryptoPriceFromDownstream(cryptoName)
	if err != nil {
		return models.CryptoPriceServiceResponse{}, err
	}
	cps.updatePriceInCache(ctx, cryptoLivePrice)
	return cryptoLivePrice, nil
}

func (cps *CryptoPriceServiceV2) cryptoPriceFromDownstream(cryptoName string) (models.CryptoPriceServiceResponse, error) {
	var cryptoPrice models.Crypto
	var err error
	switch cryptoName {
	case constants.BITCOIN_IDENTIFIER:
		cryptoPrice, err = cps.coinDeskClient.GetBTCCurrentPrice()
	case constants.ETHEREUM_IDENTIFIER:
		cryptoPrice, err = cps.cryptonatorClient.GetETHCurrentPrice()
	default:
		cryptoPrice, err = cps.coinDeskClient.GetBTCCurrentPrice()
	}
	if err != nil {
		return models.CryptoPriceServiceResponse{}, err
	}
	cryptoPriceConverted := cryptoPriceToAPIResponse(cryptoPrice)
	return cryptoPriceConverted, nil
}

func (cps *CryptoPriceServiceV2) cryptoPriceFromCache(ctx context.Context, cryptoName string) (models.CryptoPriceServiceResponse, error) {
	crypto, err := cps.storageClient.GetCryptoPrice(ctx, cryptoName)
	if err != nil {
		return models.CryptoPriceServiceResponse{}, err
	}
	return newCryptoPriceServiceResponse(crypto), nil
}

func (cps *CryptoPriceServiceV2) updatePriceInCache(ctx context.Context, cpsr models.CryptoPriceServiceResponse) {
	err := cps.storageClient.SetCryptoPrice(ctx, cryptoFromCryptoPriceServiceResponse(cpsr))
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
