package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/st3v/translator/microsoft"
)

func HandleTranslate(msg string) string {
	translator := microsoft.NewTranslator(os.Getenv("MSFT"))

	output := make(chan string)
	gochan := make(chan string)

	go func(t string, gochan chan string) {
		translation, err := translator.Translate(t, "en", "ru")
		translation, err = translator.Translate(translation, "ru", "yue")
		translation, err = translator.Translate(translation, "yue", "es")
		translation, err = translator.Translate(translation, "es", "vi")

		if err != nil {
			log.Panicf("Error during translation: %s", err.Error())
		}

		gochan <- translation
	}(msg, gochan)

	go func(gochan chan string) {
		msg := <-gochan
		translation, err := translator.Translate(msg, "vi", "en")

		if err != nil {
			log.Panicf("error at translating: %v", err)
		}

		output <- translation
	}(gochan)

	return <-output
}

func HandleDadJokes() string {
	c := &http.Client{}
	var body DadJokesResponse

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	if err != nil {
		log.Panicf("Error during getting a dad joke: %s", err.Error())
	}

	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Panic(err)
	}

	fmt.Println("Dad joke is:", body.Joke)

	defer resp.Body.Close()
	return body.Joke
}
