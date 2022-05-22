package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies/types"
	"interphlix/lib/variables"
	"net/http"
)


func GetTypes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(types.Types))
}