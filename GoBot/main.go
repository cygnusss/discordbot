package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Token is my key
const Token string = ""

// BotID is BotID
var BotID string

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend("412507247697068032", "pong")
	}

}

func main() {

	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		log.Panic(err)
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		log.Panic(err)
	}

	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("Bot is running!")

	<-make(chan struct{})
	return

}
