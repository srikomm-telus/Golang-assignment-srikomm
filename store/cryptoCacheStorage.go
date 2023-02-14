package store

import (
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	"errors"
	"go.uber.org/zap"
	"time"
)

type CryptoCacheStorage struct {
	storage *RedisClient
}

var logger, _ = zap.NewProduction()

func NewCryptoCacheStorage(client *RedisClient) CryptoCacheStorage {
	return CryptoCacheStorage{
		storage: client,
	}
}

func (c CryptoCacheStorage) SetCryptoPrice(crypto models.Crypto) (bool, error) {

	err := c.storage.SetValue(constants.US_CRYPTO_PRICE_CACHE_KEY+crypto.GetCryptoName(), crypto.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER),
		time.Duration(constants.EXPIRY_SECONDS)*time.Second)
	if err != nil {
		logger.Error(err.Error(), zap.String("crypto", crypto.GetCryptoName()))
		return false, errors.New("unable to set crypto price in cache")
	}
	logger.Info(constants.SET_CACHE_VALUE, zap.String(constants.US_CRYPTO_PRICE_CACHE_KEY+crypto.GetCryptoName(), crypto.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER)))
	err = c.storage.SetValue(constants.EU_CRYPTO_PRICE_CACHE_KEY+crypto.GetCryptoName(), crypto.GetPriceInCurrency(constants.EUR_CURRENCY_IDENTIFIER),
		time.Duration(constants.EXPIRY_SECONDS)*time.Second)
	if err != nil {
		logger.Error(err.Error(), zap.String("crypto", crypto.GetCryptoName()))
		return false, errors.New("unable to set crypto price in cache")
	}
	logger.Info(constants.SET_CACHE_VALUE, zap.String(constants.EU_CRYPTO_PRICE_CACHE_KEY+crypto.GetCryptoName(), crypto.GetPriceInCurrency(constants.EUR_CURRENCY_IDENTIFIER)))

	return true, nil
}

func (c CryptoCacheStorage) GetCryptoPrice(cryptoName string) (models.Crypto, error) {
	usdRate, err := c.storage.GetValue(constants.US_CRYPTO_PRICE_CACHE_KEY + cryptoName)
	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, errors.New("unable to fetch from cache")
	}
	eurRate, err := c.storage.GetValue(constants.EU_CRYPTO_PRICE_CACHE_KEY + cryptoName)
	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, errors.New("unable to fetch from cache")
	}

	if (usdRate == constants.EMPTY_STRING) || (eurRate == constants.EMPTY_STRING) {
		return models.Crypto{}, errors.New("invalid cache")
	}

	return models.NewCrypto(cryptoName, map[string]string{
		constants.USD_CURRENCY_IDENTIFIER: usdRate,
		constants.EUR_CURRENCY_IDENTIFIER: eurRate,
	}), nil

}
