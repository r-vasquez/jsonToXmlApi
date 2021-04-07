package main

import (
	"github.com/gin-gonic/gin"
	"github.com/r-vasquez/jsonToXmlApi/routes"
)

func main() {
	router := gin.Default()
	routes.InitializeRoutes(router)

	router.Run("localhost:1212")
}
