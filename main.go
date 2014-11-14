package main

import (
	"./handlers"
	"net/http"
)

func initHandlers() {
	http.HandleFunc("/", handlers.HandleStuff)
}

func main() {
	initHandlers()

	http.ListenAndServe(":8080", nil)
}
