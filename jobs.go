package main

type Job struct {
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

}
