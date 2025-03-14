package botHandler

import (
	"github.com/V1c-G4b/discord_bot_go/bot"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShutDownBotHandler(ctx *gin.Context) {
	if bot.StatusBot() {
		err := bot.StopBot()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao desligar o bot"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET O bot foi desligado",
		})
	}
}
