package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tanishqahluwalia25/portfolio-backend/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	//add routes here\
	routes.QueryRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "server is running")
	})

	router.Run(":" + port)

}
