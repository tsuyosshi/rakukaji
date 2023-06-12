package main

import (
	"net/http"
	"rakukaji/router"
)

func main() {
	http.HandleFunc("/", router.Hello)
	http.HandleFunc("/userCreateTest", router.UserCreateTest)
	http.ListenAndServe(":8080", nil)
}
