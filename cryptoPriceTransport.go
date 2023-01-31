package main

import (
	. "Golang-assignment-srikomm/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddHttpTransport(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/price", getCryptoPriceHandler)
	routerGroup.GET("/price/:cryptoName", getSpecificCryptoPriceHandler)
}

var logger, _ = zap.NewProduction()

func getCryptoPriceHandler(c *gin.Context) {
	cryptoPrice, err := GetCryptoPrice(DEFAULT_CRYPTO)

	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	logger.Info(PRICE_FETCH_SUCCESSFUL_MESSAGE,
		zap.String(PRICE, cryptoPrice.GetPriceInCurrency(USD_CURRENCY_IDENTIFIER)))
	c.JSON(http.StatusOK, gin.H{
		DATA: cryptoPrice,
	})
	return
}

func getSpecificCryptoPriceHandler(c *gin.Context) {
	cryptoName := c.Param(CRYPTO_NAME_PARAM_KEY)

	CryptoPrice, err := GetCryptoPrice(cryptoName)

	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	logger.Info(PRICE_FETCH_SUCCESSFUL_MESSAGE,
		zap.String(PRICE, CryptoPrice.GetPriceInCurrency(USD_CURRENCY_IDENTIFIER)))
	c.JSON(http.StatusOK, gin.H{
		DATA: CryptoPrice,
	})
	return
}
