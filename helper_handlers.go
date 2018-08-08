package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HandleDadJokes handles all !dadjoke event in messageHandler (main.go)
func HandleDadJokes(start time.Time) (string, error) {
	c := &http.Client{}
	var body DadJokesResponse

	// Create new GET request to dad jokes API
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Println("Error during getting a dad joke:\n", err.Error())
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	// Send the request and parse the response body
	resp, err := c.Do(req)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		fmt.Println("Error while decoding request body:\n", err)
		return "", err
	}

	defer req.Body.Close()
	defer resp.Body.Close()
	defer fmt.Println("Hadnled !dadjoke in:", time.Now().Sub(start))

	return body.Joke, nil
}
