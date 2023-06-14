package transponster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (io IO) logOutgoing(code int) {
	elapsed := time.Since(io.start).Milliseconds()
	log.Printf("OUTGOING[%v]: %v took %vms\n", io.id, code, elapsed)
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

	io.logOutgoing(http.StatusOK)
}

func (io IO) Image(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		io.E500(fmt.Errorf("os read of file failed: %v", err))
		return
	}

	io.W.Header().Set("Content-Type", "application/octet-stream")
	_, err = io.W.Write(b)
	if err != nil {
		io.E500(fmt.Errorf("response writing failed: %v", err))
		return
	}

	io.logOutgoing(http.StatusOK)
}

func (io IO) E400(err error, msg string) {
	if msg == "" {
		msg = "Bad Request"
	}

	code := http.StatusBadRequest
	io.logOutgoing(code)
	log.Println(msg)
	log.Println(err)
	http.Error(io.W, msg, code)
}

func (io IO) E401() {
	code := http.StatusUnauthorized
	io.logOutgoing(code)
	http.Error(io.W, "Unauthenticated", code)
}

func (io IO) E403() {
	code := http.StatusForbidden
	io.logOutgoing(code)
	http.Error(io.W, "Unauthorised", code)
}

func (io IO) E404() {
	code := http.StatusNotFound
	io.logOutgoing(code)
	http.Error(io.W, "Not Found", code)
}

func (io IO) E500(e error) {
	code := http.StatusInternalServerError
	io.logOutgoing(code)
	log.Println(e)
	http.Error(io.W, "Internal Server Error", code)
}

func (io IO) E501() {
	code := http.StatusNotImplemented
	io.logOutgoing(code)
	http.Error(io.W, "Not Yet Implemented", code)
}
