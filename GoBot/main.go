package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/st3v/translator/microsoft"
)

// BotID is BotID
var BotID string

func handleTranslate(t string) string {
	translator := microsoft.NewTranslator(os.Getenv("MSFT"))

	translation, err := translator.Translate(t, "en", "ru")
	translation, err = translator.Translate(translation, "ru", "yue")
	translation, err = translator.Translate(translation, "yue", "vi")
	translation, err = translator.Translate(translation, "vi", "es")
	translation, err = translator.Translate(translation, "es", "en")

	if err != nil {
		log.Panicf("Error during translation: %s", err.Error())
	}
	fmt.Printf("TRANSLATION IS: %s", translation)
	return translation
}

func handleDadJokes(ch chan string) <-chan string {

	c := &http.Client{}
	var b DadJokesResponse

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)

	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()
	ch <- b.Joke
	return ch

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	transBool := strings.HasPrefix(m.Content, "/translate")
	dadjBool := strings.HasPrefix(m.Content, "/dadjoke")
	helpBool := strings.HasPrefix(m.Content, "/donkey")

	if helpBool {
		_, _ = s.ChannelMessageSend("412498257655365633", "'/translate' - translates whatever is after the tag\n'/dadjoke' - get a random dad joke\n¯\\_(ツ)_/¯")
	}

	if transBool {
		_, _ = s.ChannelMessageSend("412498257655365633", handleTranslate(m.Content[11:]))
	}

	if dadjBool {
		ch := make(chan string)
		go handleDadJokes(ch)
		_, _ = s.ChannelMessageSend("412498257655365633", <-ch)
	}

	if strings.Contains(strings.ToLower(m.Content), "bitcoin") {
		_, _ = s.ChannelMessageSend("412498257655365633", "BITCOIN IS A BUBBLE")
	}

	if strings.ToLower(m.Content) == "david" {
		_, _ = s.ChannelMessageSend("412498257655365633", "https://giphy.com/gifs/3djolNOedd5pS")
	}

}

func main() {

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
