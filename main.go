package main

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
)

// BotID is BotID
var BotID string

func init() {
	gotenv.Load()
}

func main() {
	go StartServer()

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

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

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	transBool := strings.HasPrefix(m.Content, "/translate")
	dadjBool := strings.HasPrefix(m.Content, "/dadjoke")
	helpBool := strings.HasPrefix(m.Content, "/donkey")

	c := m.ChannelID

	if helpBool {
		_, _ = s.ChannelMessageSend(c, "'/translate' - translates whatever is after the tag into donkey language\n'/dadjoke' - get a random dad joke\n¯\\_(ツ)_/¯")
	}

	if transBool {
		if len(m.Content) <= 11 {
			_, _ = s.ChannelMessageSend(c, "Message is too short, try again!")
		} else {
			msg := HandleTranslate(m.Content[11:])

			_, _ = s.ChannelMessageSend(c, msg)
		}
	}

	if dadjBool {
		joke := HandleDadJokes()
		_, _ = s.ChannelMessageSend(c, joke)
	}

	if strings.Contains(strings.ToLower(m.Content), "bitcoin") {
		_, _ = s.ChannelMessageSend(c, "BITCOIN IS A BUBBLE")
	}

	if strings.ToLower(m.Content) == "david" {
		_, _ = s.ChannelMessageSend(c, "https://giphy.com/gifs/3djolNOedd5pS")
	}
}
