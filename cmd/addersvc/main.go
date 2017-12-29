package main

import (
	"net/http"

	"github.com/dandeliondeathray/goosedev/adder"
)

func main() {
	router := adder.Router()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
