package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type dResp struct {
	Joke string `json:"joke"`
}

func donkeyHandler(w http.ResponseWriter, r *http.Request) {
	resp := dResp{"hello world"}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func StartServer() {

	http.HandleFunc("/donkey", donkeyHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8081"
	}

	http.ListenAndServe(port, nil)

}
