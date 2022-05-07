package main

import (
	"interphlix/lib/handler"
	"net/http"
)

var (
	PORT = ":8000"
)

func main() {

	go handler.Main()

	http.ListenAndServe(PORT, handler.Router)
}