package main

import (
	"net/http"

	"github.com/bradlsum/stockey/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe("localhost:8081", mux)
}