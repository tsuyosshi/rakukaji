package router

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
