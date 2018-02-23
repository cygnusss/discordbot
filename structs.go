package main

type DResp struct {
	Joke string `json:"joke"`
}

// Stringify converts json into a go struct
type DadJokesResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}
