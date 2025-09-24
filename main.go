package main

import (
	"github.com/gin-gonic/gin"
	"github.com/melospot/openzam-api/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":5000")
}
