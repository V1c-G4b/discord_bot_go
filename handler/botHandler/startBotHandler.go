package botHandler

import (
	"github.com/V1c-G4b/discord_bot_go/bot"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartBotHandler(ctx *gin.Context, token string) {
	if !bot.StatusBot() {
		err := bot.StartBot(token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao ligar o bot"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET O bot foi ligado",
		})
	}
}
