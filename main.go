package main

import (
	"github.com/V1c-G4b/discord_bot_go/bot"
	"github.com/V1c-G4b/discord_bot_go/config"
	"github.com/V1c-G4b/discord_bot_go/router"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

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
