package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Stringify converts json into a go struct
type Stringify struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// Token is my key
const Token string = "NDEyNDk3NjE1MzgxMzMxOTY5.DWLJGA.gbWEk7b4a6F-7Ik2BPdCfda4bS4"

// BotID is BotID
var BotID string

func handleDadJokes() string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	var record Stringify

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Panic(err)
	}

	log.Println(record.Joke)

	defer resp.Body.Close()

	return record.Joke

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if strings.Contains(m.Content, "bitcoin") {
		_, _ = s.ChannelMessageSend("412498257655365633", "BITCOIN IS A BUBBLE")
	}

	if strings.Contains(m.Content, "dad joke") {
		_, _ = s.ChannelMessageSend("412498257655365633", handleDadJokes())
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend("412507247697068032", "https://giphy.com/gifs/3djolNOedd5pS")
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
