package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Job struct {
	w    http.ResponseWriter
	r    *http.Request
	Done chan bool
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(wp chan chan Job) Worker {
	return Worker{
		WorkerPool: wp,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
	log.Println("worker started")

	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case j := <-w.JobChannel:
				j.Run()
			case <-w.quit:
				return
			}
		}
	}()
}

func (j *Job) Run() {
	var resp DResp

	log.Println("Worker is running")

	resp.Joke = HandleDadJokes()

	respJSON, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	j.w.Header().Set("Content-Type", "application/json")
	j.w.WriteHeader(http.StatusOK)
	j.w.Write(respJSON)

	j.Done <- true
}
