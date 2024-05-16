package main

import (
	"log"
	"net/http"
	"webform-golang/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal(err)
	}
}
