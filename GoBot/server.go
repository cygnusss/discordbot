package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	newrelic "github.com/newrelic/go-agent"
)

type dResp struct {
	Joke string `json:"joke"`
}

func dadJoke(w http.ResponseWriter, r *http.Request) {
	var resp dResp

	resp.Joke = HandleDadJokes()

	respJSON, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func StartServer() {

	config := newrelic.NewConfig("DiscordBot", os.Getenv("NEWRELIC"))
	App, err := newrelic.NewApplication(config)

	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(App, "/dadjoke", dadJoke))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8081"
	}

	http.ListenAndServe(port, nil)

}
