package main

import (
	"encoding/json"
	"net/http"
)

// JobQueue is a channel of Jobs, Job struct can be found in the structs file
var JobQueue chan Job

func NewWorker(wp chan chan Job) Worker {
	return Worker{
		WorkerPool: wp,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
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
