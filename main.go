package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
)

// BotID globally stores the bot's ID
var BotID string

func init() { gotenv.Load() }

func main() {
	// Create discord session
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Error while creating discord session:\n", err)
		return
	}

	// Get user info for the bot
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("Error while getting user info:\n", err)
		return
	}
	// Get the id of the bot
	BotID = u.ID

	dg.AddHandler(messageHandler)

	// Open the websocket and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error while openning the websocket:\n", err)
		return
	}

	fmt.Println("Discord Bot is running!")

	defer dg.Close()
	<-make(chan struct{})
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	start := time.Now()
	// If message is coming from the bot do nothing
	if m.Author.ID == BotID {
		return
	}

	dadjBool := strings.HasPrefix(m.Content, "!dadjoke")
	helpBool := strings.HasPrefix(m.Content, "!donkey")

	// c is the ID of the message's channel
	c := m.ChannelID

	if helpBool {
		_, _ = s.ChannelMessageSend(c, "'!dadjoke' - get a random dad joke\n¯\\_(ツ)_/¯")
	}

	// Handles all !dadjoke events
	if dadjBool {
		joke, err := HandleDadJokes(start)
		if err != nil {
			fmt.Println("Error while sending a request to dad-jokes API:\n", err)
		} else {
			_, _ = s.ChannelMessageSend(c, joke)
		}
	}
}
