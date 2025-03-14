package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter(discordToken string) {
	router := gin.Default()

	initializeRoutes(router, discordToken)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
