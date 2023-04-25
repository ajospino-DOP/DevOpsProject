package server

import "net/http"

func routes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/info", info)
}