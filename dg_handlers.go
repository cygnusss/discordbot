package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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
