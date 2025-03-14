package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"sync"
)

var (
	discordSession *discordgo.Session
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

func StartBot(discordToken string) error {
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

func BotGuilds() []*discordgo.Guild {
	botMutex.Lock()
	defer botMutex.Unlock()

	if discordSession != nil {
		return discordSession.State.Guilds
	}
	return nil
}
