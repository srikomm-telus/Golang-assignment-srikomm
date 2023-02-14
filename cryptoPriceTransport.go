package main

import (
	"Golang-assignment-srikomm/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddHttpTransport(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/price", getCryptoPriceHandler)
	routerGroup.GET("/price/:cryptoName", getSpecificCryptoPriceHandler)
}

var logger, _ = zap.NewProduction()
var cps CryptoPriceService

func getCryptoPriceHandler(c *gin.Context) {
	var err error
	cps, err = NewCryptoPriceService(constants.DEFAULT_CRYPTO, c)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	cryptoPrice, err := cps.GetCryptoPrice(constants.DEFAULT_CRYPTO)

	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			constants.ERROR: err.Error(),
		})
		return
	}

	logger.Info("Crypto Price fetched successfully",
		zap.String(constants.PRICE, cryptoPrice.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER)))
	c.JSON(http.StatusOK, gin.H{
		constants.DATA: cryptoPrice,
	})
	return
}

func getSpecificCryptoPriceHandler(c *gin.Context) {
	var err error
	cryptoName := c.Param(constants.CRYPTO_NAME_PARAM_KEY)
	cps, err = NewCryptoPriceService(cryptoName, c)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	CryptoPrice, err := cps.GetCryptoPrice(cryptoName)

	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			constants.ERROR: err.Error(),
		})
		return
	}

	logger.Info("Crypto Price fetched successfully",
		zap.String(constants.PRICE, CryptoPrice.GetPriceInCurrency(constants.USD_CURRENCY_IDENTIFIER)))
	c.JSON(http.StatusOK, gin.H{
		constants.DATA: CryptoPrice,
	})
	return
}
