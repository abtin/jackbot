package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		fmt.Println("No token provided. Set DISCORD_BOT_TOKEN environment variable and re-run the bot.")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		os.Exit(1)
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		os.Exit(1)
	}
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("jackbot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// messageCreate is registered as a handler and responses to the input messaged from the users
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// A Simple, stupid way of providing the services. Good for 1st iteration ;)
	switch strings.ToLower(m.Content) {
	case "time":
		now := time.Now()
		_, err := s.ChannelMessageSend(m.ChannelID, now.Format("Mon Jan 2 15:04:05 MST 2006"))
		if err != nil {
			fmt.Println(err)
		}
	}
}
