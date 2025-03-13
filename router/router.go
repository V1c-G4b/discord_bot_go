package router

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(discordSession *discordgo.Session) {
	router := gin.Default()

	initializeRoutes(router, discordSession)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
