package main

import (
	. "Golang-assignment-srikomm/constants"
	. "Golang-assignment-srikomm/factory"
	. "Golang-assignment-srikomm/models"
	. "Golang-assignment-srikomm/redis"
	. "Golang-assignment-srikomm/util"
	"errors"
	"go.uber.org/zap"
	"time"
)

func GetCryptoPrice(cryptoName string) (CryptoPriceServiceResponse, error) {

	StoredCryptoPrice, err := cryptoPriceFromCache(cryptoName)
	if err == nil {
		return StoredCryptoPrice, err
	}

	CryptoLivePrice, err := cryptoPriceDownstream(cryptoName)
	if err != nil {
		return CryptoPriceServiceResponse{}, err
	}

	updatePriceInCache(CryptoLivePrice, cryptoName)

	return CryptoLivePrice, nil
}

func cryptoPriceDownstream(cryptoName string) (CryptoPriceServiceResponse, error) {
	cryptoClient := CryptoClientFactory(cryptoName)

	cryptoPrice, err := cryptoClient.GetCurrentPrice()

	if err != nil {
		return CryptoPriceServiceResponse{}, err
	}

	cryptoPriceConverted := CryptoPriceToAPIResponse(cryptoPrice)

	return cryptoPriceConverted, nil
}

func cryptoPriceFromCache(cryptoName string) (CryptoPriceServiceResponse, error) {
	rdb := RedisClient{}
	usdRate := rdb.GetValue(US_CRYPTO_PRICE_CACHE_KEY + cryptoName)
	eurRate := rdb.GetValue(EU_CRYPTO_PRICE_CACHE_KEY + cryptoName)

	if (usdRate == EMPTY_STRING) || (eurRate == EMPTY_STRING) {
		return CryptoPriceServiceResponse{}, errors.New(INVALID_CACHE_ERROR_MESSAGE)
	}

	return NewCryptoPriceServiceResponse(usdRate, eurRate, PRICE_IS_FROM_CACHE, cryptoName), nil
}

func updatePriceInCache(cpsr CryptoPriceServiceResponse, cryptoName string) {
	rdb := RedisClient{}
	rdb.SetValue(US_CRYPTO_PRICE_CACHE_KEY+cryptoName, cpsr.Data[USD_CURRENCY_IDENTIFIER],
		time.Duration(EXPIRY_SECONDS)*time.Second)
	logger.Info(SET_CACHE_VALUE, zap.String(US_CRYPTO_PRICE_CACHE_KEY+cryptoName, cpsr.Data[USD_CURRENCY_IDENTIFIER]))
	rdb.SetValue(EU_CRYPTO_PRICE_CACHE_KEY+cryptoName, cpsr.Data[EUR_CURRENCY_IDENTIFIER],
		time.Duration(EXPIRY_SECONDS)*time.Second)
	logger.Info(SET_CACHE_VALUE, zap.String(EU_CRYPTO_PRICE_CACHE_KEY+cryptoName, cpsr.Data[EUR_CURRENCY_IDENTIFIER]))
}
