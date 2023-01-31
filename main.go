package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routerGroup := router.Group("/crypto")
	AddHttpTransport(routerGroup)
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Failed to start service on localhost 8080 ", err)
		return
	}
}
