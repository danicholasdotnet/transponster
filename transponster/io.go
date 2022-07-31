package transponster

import (
	"fmt"
	"net/http"
)

type IO struct {
	R *http.Request
	W http.ResponseWriter
}

type Detail struct {
	method string
	url    string
	ua     string
	src    string
}

func GetDetail(r *http.Request) string {
	return fmt.Sprintf("%+v", Detail{
		method: r.Method,
		url:    r.URL.Path,
		ua:     r.UserAgent(),
		src:    r.RemoteAddr,
	})
}
