package main

import "net/http"

type DResp struct {
	Joke string `json:"joke"`
}

type DadJokesResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type Job struct {
	w    http.ResponseWriter
	r    *http.Request
	Done chan bool
}

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}
