package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies/types"
	"interphlix/lib/variables"
	"net/http"
)


func GetTypes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	data, status := types.GetTypes()
	res.WriteHeader(status)
	res.Write(data)
}