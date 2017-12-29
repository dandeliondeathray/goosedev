package subtractor

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dandeliondeathray/goosedev/mylib"
	"github.com/gorilla/mux"
)

// Router returns a Gorilla Mux HTTP router for the subtractor service.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/subtract/{x:[0-9]+}/{y:[0-9]+}", doSubtract)
	return r
}

// doSubtract handles HTTP requests to subtract two numbers
func doSubtract(w http.ResponseWriter, req *http.Request) {
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

	subtractor := mylib.NewSubtractor(x)
	result := subtractor.Op(y)

	body := fmt.Sprintf("%d", result)
	w.Write([]byte(body))
}
