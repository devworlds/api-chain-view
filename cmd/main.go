package main

import (
	"net/http"

	"github.com/devworlds/api-chain-view/internal/ethereum"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	r.GET("/create-wallet", func(c *gin.Context) {
		wallet, err := ethereum.GenerateWallet()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Wallet": wallet,
		})
	})
	r.Run()
}
