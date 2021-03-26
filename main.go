package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flyingpot/codewars-readme-stats/api"
)

func main() {
	http.Handle("/", http.HandlerFunc(api.Handler))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
