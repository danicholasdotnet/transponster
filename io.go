package transponster

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type IO struct {
	W     http.ResponseWriter
	R     *http.Request
	id    int
	start time.Time
}

func NewIO(w http.ResponseWriter, r *http.Request) IO {
	io := IO{
		W:     w,
		R:     r,
		id:    rand.Intn(9999),
		start: time.Now(),
	}

	io.logIncoming()

	return io
}

type detail struct {
	method string
	url    string
	ua     string
	src    string
}

func getDetail(r *http.Request) string {
	return fmt.Sprintf("%+v", detail{
		method: r.Method,
		url:    r.URL.Path,
		ua:     r.UserAgent(),
		src:    r.RemoteAddr,
	})
}
