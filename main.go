package main

import (
	"fmt"
	"github.com/V1c-G4b/discord_bot_go/router"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	discordSession *discordgo.Session
	discordToken   string
	botMutex       sync.Mutex
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			return
		}
	}
}

func StartBot() error {
	botMutex.Lock()
	defer botMutex.Unlock()

	if discordSession != nil {
		return fmt.Errorf("o bot já está em execução")
	}

	s, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		return err
	}
	s.AddHandler(messageCreate)

	if err = s.Open(); err != nil {
		return err
	}
	discordSession = s
	return nil
}

func StopBot() error {
	botMutex.Lock()
	defer botMutex.Unlock()

	if discordSession == nil {
		return fmt.Errorf("o bot não está em execução")
	}
	err := discordSession.Close()
	discordSession = nil
	return err
}

func StatusBot() bool {
	botMutex.Lock()
	defer botMutex.Unlock()

	return discordSession != nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado")
	}

	discordToken = os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Println("No Discord token found")
		return
	}

	if err := StartBot(); err != nil {
		log.Fatalf("Erro ao iniciar o bot: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		router.InitializeRouter(discordSession)
	}()

	<-stop
	log.Println("Encerrando aplicação...")
	if StatusBot() {
		StopBot()
	}
}
