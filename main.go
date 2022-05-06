package main

import (
	"encoding/json"
	"fmt"
	"interphlix/lib/crawler/tinyzone/movies"
	"interphlix/lib/crawler/tinyzone/tvshows"
	"interphlix/lib/crawler/tinyzone"
	"interphlix/lib/types"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	tinyzone.Main()

	http.HandleFunc("/", Stats)
	http.HandleFunc("/movies", Movies)
	http.HandleFunc("/tv-shows", TvShows)

	http.ListenAndServe(":8000", nil)
}

func Stats(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	data := fmt.Sprintf(`[{"total": "%d", "position": "%d", "available": "%d"},{"total": "%d", "position": "%d", "available": "%d"}]`, len(movies.Movies), movies.Position, movies.Available, len(tvshows.TvShows), tvshows.Position, tvshows.Available)
	res.Write([]byte(data))
}

func Movies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var Movies []types.Movie
	var Response []types.Movie
	Start := req.URL.Query().Get("start")
	End := req.URL.Query().Get("end")

	data, _ := ioutil.ReadFile("./DB/movies.json")
	json.Unmarshal(data, &Movies)
	if End == "0" && (Start == "" || Start == "0"){
		data := types.JsonMarshal(Movies)
		res.Write(data)
		return
	}else if End == "0" {
		start, err := strconv.Atoi(Start)
		if err != nil {
			res.Write([]byte(`{"error": "invalid start number"}`))
		}
		Response = Movies[start:]
		data := types.JsonMarshal(Response)
		res.Write(data)
		return
	}else if End == "" {
		start, err := strconv.Atoi(Start)
		if err != nil {
			res.Write([]byte(`{"error": "invalid start number"}`))
		}
		Response = Movies[start:start+20]
		data := types.JsonMarshal(Response)
		res.Write(data)
		return
	}else {
		start, err := strconv.Atoi(Start)
		if err != nil {
			res.Write([]byte(`{"error": "invalid start number"}`))
		}
		end, err := strconv.Atoi(End)
		if err != nil {
			res.Write([]byte(`{"error": "invalid end number"}`))
		}
		Response = Movies[start:end]
		data := types.JsonMarshal(Response)
		res.Write(data)
		return
	}
}

func TvShows(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var TvShows []types.Movie
	var Response []types.Movie
	Start := req.URL.Query().Get("start")
	End := req.URL.Query().Get("end")

	data, _ := ioutil.ReadFile("./DB/tvshows.json")
	json.Unmarshal(data, &TvShows)
	if End == "0" && (Start == "" || Start == "0"){
		data := types.JsonMarshal(Movies)
		res.Write(data)
		return
	}else if End == "0" {
		start, err := strconv.Atoi(Start)
		if err != nil {
			res.Write([]byte(`{"error": "invalid start number"}`))
		}
		Response = TvShows[start:]
		data := types.JsonMarshal(Response)
		res.Write(data)
		return
	}else if End == "" {
		start, err := strconv.Atoi(Start)
		if err != nil {
			res.Write([]byte(`{"error": "invalid start number"}`))
		}
		Response = TvShows[start:start+20]
		data := types.JsonMarshal(Response)
		res.Write(data)
		return
	}else {
		start, err := strconv.Atoi(Start)
		if err != nil {
			res.Write([]byte(`{"error": "invalid start number"}`))
		}
		end, err := strconv.Atoi(End)
		if err != nil {
			res.Write([]byte(`{"error": "invalid end number"}`))
		}
		Response = TvShows[start:end]
		data := types.JsonMarshal(Response)
		res.Write(data)
		return
	}
}