package main

import (
	"Golang-assignment-srikomm/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddHttpTransport(routerGroup *gin.RouterGroup, cps CryptoPriceService, cpsv2 CryptoPriceServiceV2) {
	routerGroup.GET("/price", getBTCPriceHandler(cps))
	routerGroup.GET("/price/BTC", getBTCPriceHandler(cps))
	routerGroup.GET("/price/v2/:cryptoName", getSpecificCryptoPriceHandler(cpsv2))
}

var logger, _ = zap.NewProduction()

func getBTCPriceHandler(cps CryptoPriceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		CryptoPrice, err := cps.GetCryptoPrice(constants.BITCOIN_IDENTIFIER)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
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
}

func getSpecificCryptoPriceHandler(cpsv2 CryptoPriceServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		cryptoName := c.Param(constants.CRYPTO_NAME_PARAM_KEY)
		CryptoPrice, err := cpsv2.GetCryptoPrice(cryptoName)

		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
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
}
