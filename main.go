package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/my/repo/src/routes"

	// init database connection
	_ "github.com/my/repo/src/database"
)

var router *gin.Engine

func main() {
	fmt.Println(`Hello, Pornixen`)
	initRoutes()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initRoutes() {
	router = gin.Default()
	router.Use(CORSMiddleware())
	routes.InitChannelRoutes(router)
	router.Run()
}
