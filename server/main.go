package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type data struct {
	Client string `json:"client"`
	Message  string    `json:"message"`
}

var m *data = &data{
	Client: "Shell",
	Message:  "Hi from Dave!",
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		j, _ := json.Marshal(m)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func main() {
	http.HandleFunc("/shell", handler)
	log.Println("server running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
