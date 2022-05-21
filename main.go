package main

import (
	"apiaive/api"
	"log"
	"net/http"
)

// @title APIAIVE api documentation
// @version 1.0.0

//@host localhost:3000
//@BasePath /api
func main() {

	err := http.ListenAndServe(":3000", api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
