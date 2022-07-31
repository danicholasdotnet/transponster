package transponster

import (
	"fmt"
	"net/http"
)

type IO struct {
	R *http.Request
	W http.ResponseWriter
}

func Details(r *http.Request) string {
	return fmt.Sprintf(r.Method, r.URL, r.UserAgent(), r.RemoteAddr)
}
