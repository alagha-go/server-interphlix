package movies

import (
	"encoding/json"
	// "interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
)


func UploadMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	// err, status := accounts.ValidateRequest(req, "user")
	// if err != nil {
	// 	res.WriteHeader(status)
	// 	res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
	// 	return
	// }
	var Movie movies.Movie
	err := json.NewDecoder(req.Body).Decode(&Movie)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid json"}))
		return
	}
	data, status := movies.UploadOneMovie(Movie)
	res.WriteHeader(status)
	res.Write(data)
}