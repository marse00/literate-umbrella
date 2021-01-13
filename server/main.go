package main

import (
	"log"
	"middleware"
	"net/http"
	"router"
)

func main() {

	err := middleware.OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}
	r := router.CreateRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
