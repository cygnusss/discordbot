package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	newrelic "github.com/newrelic/go-agent"
)

func StartServer() {
	config := newrelic.NewConfig("DiscordBot", os.Getenv("NEWRELIC"))
	app, err := newrelic.NewApplication(config)

	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/dadjoke", func(w http.ResponseWriter, r *http.Request) {
		j := Job{w, r, make(chan bool)}
		JobQueue <- j
		for {
			select {
			case <-j.Done:
				return
			}
		}
	}))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8081"
	}

	fmt.Println("server is running")
	http.ListenAndServe(port, nil)
}
