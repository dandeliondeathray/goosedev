package adder_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/dandeliondeathray/goosedev/adder"
)

func TestAdder_2Plus1_Is3(t *testing.T) {
	request := httptest.NewRequest("GET", "/add/2/1", nil)
	recorder := httptest.NewRecorder()

	router := adder.Router()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected response code 200 OK but got %d", recorder.Code)
	}

	body, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal("Couldn't read body, because ", err)
	}

	result, err := strconv.Atoi(string(body))
	if err != nil {
		t.Fatalf("Result was not an int: '%s'", string(body))
	}
	if result != 3 {
		t.Fatalf("Expected result 3 but got %d", result)
	}
}
func TestAdder_WrongURL_StatusIs404(t *testing.T) {
	request := httptest.NewRequest("GET", "/notcorrecturl/2/1", nil)
	recorder := httptest.NewRecorder()

	router := adder.Router()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("Expected response code 404 OK but got %d", recorder.Code)
	}
}
