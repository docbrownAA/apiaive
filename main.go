package main

import (
	"gduvinage/api"
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":3000", api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
