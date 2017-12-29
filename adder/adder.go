package adder

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dandeliondeathray/goosedev/mylib"
	"github.com/gorilla/mux"
)

// Router returns a Gorilla Mux HTTP router for the adder service.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/add/{x:[0-9]+}/{y:[0-9]+}", doAdd)
	return r
}

// doAdd handles HTTP requests to add two numbers
func doAdd(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	stringX := vars["x"]
	stringY := vars["y"]

	x, err := strconv.Atoi(stringX)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(stringY)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	adder := mylib.NewAdder(x)
	result := adder.Op(y)

	body := fmt.Sprintf("%d", result)
	w.Write([]byte(body))
}
