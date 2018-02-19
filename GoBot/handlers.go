package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/st3v/translator/microsoft"
)

func HandleTranslate(t string) string {
	translator := microsoft.NewTranslator(os.Getenv("MSFT"))

	translation, err := translator.Translate(t, "en", "ru")
	translation, err = translator.Translate(translation, "ru", "yue")
	translation, err = translator.Translate(translation, "yue", "vi")
	translation, err = translator.Translate(translation, "vi", "es")
	translation, err = translator.Translate(translation, "es", "en")

	if err != nil {
		log.Panicf("Error during translation: %s", err.Error())
	}

	return translation
}

func HandleDadJokes(ch chan string) <-chan string {

	c := &http.Client{}
	var b DadJokesResponse

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	if err != nil {
		log.Panicf("Error during getting a dad joke: %s", err.Error())
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
