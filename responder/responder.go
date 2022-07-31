package responder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type IO struct {
	R *http.Request
	W http.ResponseWriter
}

func Details(r *http.Request) string {
	return fmt.Sprintf(r.Method, r.URL, r.UserAgent(), r.RemoteAddr)
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
}

func (io IO) E400(err error, msg string) {
	log.Println("400 Returned For Following Request: ", Details(io.R))
	log.Println(msg)
	log.Println(err)
	if msg == "" {
		msg = "Bad Request"
	}
	http.Error(io.W, msg, http.StatusBadRequest)
}
func (io IO) E500(e error) {
	log.Println("500 Returned For Following Request: ", Details(io.R))
	log.Println(e)
	http.Error(io.W, "Internal Server Error", http.StatusInternalServerError)
}
