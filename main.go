package main

import (
	"github.com/anupamraj16/db"
	"github.com/anupamraj16/routes"
	"github.com/gin-gonic/gin"
)

// no shorthand declaration of a variable ( using := ) outside the body of a function
var PORT string = ":5001"

func main() {
	db.InitDB()
	server := gin.Default()
	// Default returns an Engine instance
	// with the Logger and Recovery middleware already attached

	routes.RegisterRoutes(server)

	server.Run(PORT)
}
