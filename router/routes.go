package router

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine, discordSession *discordgo.Session) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", func(c *gin.Context) {
			if discordSession != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": "GET O bot está em execução",
				})
			} else {
				c.JSON(http.StatusOK, "O bot está desligado")
			}
		})
	}
}
