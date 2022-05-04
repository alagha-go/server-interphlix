package main

import (
	"fmt"
	"interphlix/lib/crawler/tinyzone/movies"
	"net/http"
)

func main() {
	movies.Main()

	http.HandleFunc("/", Stats)

	http.ListenAndServe(":8000", nil)
}

func Stats(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	data := fmt.Sprintf(`{"total": "%d", "position": "%d", "available": "%s"`, len(movies.Movies), movies.Position, movies.Available)
	res.Write([]byte(data))
}