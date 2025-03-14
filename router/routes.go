package router

import (
	"github.com/V1c-G4b/discord_bot_go/handler/botHandler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, discordToken string) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", botHandler.GetBotStatusHandler)
		v1.GET("/shutdown", botHandler.ShutDownBotHandler)
		v1.GET("/start", func(ctx *gin.Context) {
			botHandler.StartBotHandler(ctx, discordToken)
		})
	}
}
