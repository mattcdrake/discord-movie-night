package main

import (
	"fmt"
	//"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create a new Discord API session
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error when creating Discord session: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

// Responds to the MESSAGE_CREATE event and routes to the appropriate action.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
}
