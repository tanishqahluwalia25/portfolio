package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tanishqahluwalia25/portfolio-backend/controllers"
)

func QueryRoutes(router *gin.Engine) {

	router.GET("/query", controllers.MakeQuery())

}
