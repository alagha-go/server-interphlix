package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies/ratings"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMovieRatings(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	ID, err := primitive.ObjectIDFromHex(req.URL.Query().Get("id"))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	data, status := ratings.GetMovieRatings(ID)
	res.WriteHeader(status)
	res.Write(data)
}