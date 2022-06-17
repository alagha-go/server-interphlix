package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies/casts"
	"interphlix/lib/variables"
	"net/http"
	"strconv"
)


func GetAllCasts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := casts.GetAllCasts(round)
	res.WriteHeader(status)
	res.Write(data)
}