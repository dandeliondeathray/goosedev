package main

import (
	"net/http"

	"github.com/dandeliondeathray/goosedev/subtractor"
)

func main() {
	router := subtractor.Router()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
