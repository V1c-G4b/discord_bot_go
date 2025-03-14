package botHandler

import (
	"github.com/V1c-G4b/discord_bot_go/bot"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBotStatusHandler(ctx *gin.Context) {
	if bot.StatusBot() {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": bot.BotGuilds(),
		})
	} else {
		ctx.JSON(http.StatusOK, "O bot está desligado")
	}
}
