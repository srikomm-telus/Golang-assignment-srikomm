package main

import (
	"Golang-assignment-srikomm/client"
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/store"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		router            = gin.Default()
		routerGroup       = router.Group("/crypto")
		ctx               = context.Background()
		coinDeskClient    = client.NewCoinDeskClient(constants.COINDESK_ENDPOINT)
		cryptonatorClient = client.NewCryptonatorClient(constants.CRYPTONATOR_ENDPOINT)
		redisClient, err  = store.NewRedisClient(ctx, constants.PRODUCTION)
		cps               = NewCryptoPriceService(coinDeskClient, redisClient)
		cpsv2             = NewCryptoPriceServiceV2(coinDeskClient, cryptonatorClient, redisClient)
	)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	AddHttpTransport(routerGroup, cps, cpsv2)
	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Failed to start service on localhost 8080 ", err)
		return
	}
}
