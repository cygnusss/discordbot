package main

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
)

var BotID string

const (
	maxJob    = 10000
	maxWorker = 1000
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	dadjBool := strings.HasPrefix(m.Content, "/dadjoke")
	helpBool := strings.HasPrefix(m.Content, "/donkey")

	c := m.ChannelID

	if helpBool {
		_, _ = s.ChannelMessageSend(c, "'/dadjoke' - get a random dad joke\n¯\\_(ツ)_/¯")
	}

	if dadjBool {
		joke := HandleDadJokes()
		_, _ = s.ChannelMessageSend(c, joke)
	}

	if strings.ToLower(m.Content) == "david" {
		_, _ = s.ChannelMessageSend(c, "https://giphy.com/gifs/3djolNOedd5pS")
	}
}

func init() { gotenv.Load() }

func main() {
	stop := make(chan bool)
	wPool := make(chan chan Job, maxWorker)
	JobQueue = make(chan Job, maxJob)

	for i := 0; i < maxWorker; i++ {
		w := NewWorker(wPool)
		w.Start()
	}

	go func() {
		for j := range JobQueue {
			go func(j Job) {
				jChan := <-wPool
				jChan <- j
			}(j)
		}
	}()

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
	<-stop
	return
}
