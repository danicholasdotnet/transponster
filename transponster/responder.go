package transponster

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LogCodeAndRequest(code int, request *http.Request) {
	log.Println("[", code, "] --> { ", Details(request), " }")
}

func (io IO) Success(i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		io.E500(fmt.Errorf("json marshal of data failed: %v", err))
		return
	}

	io.W.Header().Set("Content-Type", "application/json")
	_, err = io.W.Write(b)
	if err != nil {
		io.E500(fmt.Errorf("response writing failed: %v", err))
		return
	}

	LogCodeAndRequest(http.StatusOK, io.R)
}

func (io IO) E400(err error, msg string) {
	if msg == "" {
		msg = "Bad Request"
	}

	code := http.StatusBadRequest
	LogCodeAndRequest(code, io.R)
	log.Println(msg)
	log.Println(err)
	http.Error(io.W, msg, code)
}

func (io IO) E404() {
	code := http.StatusNotFound
	LogCodeAndRequest(code, io.R)
	http.Error(io.W, "Not Found", code)
}

func (io IO) E500(e error) {
	code := http.StatusInternalServerError
	LogCodeAndRequest(code, io.R)
	log.Println(e)
	http.Error(io.W, "Internal Server Error", code)
}

func (io IO) E501() {
	code := http.StatusNotImplemented
	LogCodeAndRequest(code, io.R)
	http.Error(io.W, "Not Yet Implemented", code)
}
