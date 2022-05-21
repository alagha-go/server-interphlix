package movies

import (
	"interphlix/lib/movies/types"
	"interphlix/lib/variables"
	"net/http"
)


func GetTypes(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(types.Types))
}