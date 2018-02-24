package main

type DResp struct {
	Joke string `json:"joke"`
}

type DadJokesResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}
