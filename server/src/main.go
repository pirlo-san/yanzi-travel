package main

import (
	"home"
	"net/http"
)

func main() {
	http.HandleFunc("/", home.HomePage)
	http.ListenAndServe(":80", nil)
}
