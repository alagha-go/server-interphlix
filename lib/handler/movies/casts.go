package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies/casts"
	"interphlix/lib/variables"
	"net/http"
)


func GetAllCasts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	data, status := casts.GetAllCasts()
	res.WriteHeader(status)
	res.Write(data)
}