package main

import (
	"github.com/V1c-G4b/discord_bot_go/bot"
	"github.com/V1c-G4b/discord_bot_go/router"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	discordSession *discordgo.Session
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado")
	}

	var discordToken = os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Println("No Discord token found")
		return
	}

	if err := bot.StartBot(discordToken); err != nil {
		log.Fatalf("Erro ao iniciar o bot: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		router.InitializeRouter(discordToken)
	}()

	<-stop
	log.Println("Encerrando aplicação...")
	if bot.StatusBot() {
		bot.StopBot()
	}
}
