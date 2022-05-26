package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func GetMovies(index int) ([]byte, int) {
	var length int
	var Movies []Movie
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	cursor, err := collection.Find(context.Background(), bson.M{})
	variables.HandleError(err, "movies", "GetMovies", "error while getting movies from the local database")
	cursor.All(context.Background(), &Movies)
	if index == 0 {
		length = 30
	}else if index < 0{
		length = 0
	}else {
		length = (index*10)+30
	}
	if length == 0 {
		return variables.JsonMarshal(variables.Error{Error: "invalid index"}), http.StatusBadRequest
	}else if length == 30 {
		if len(Movies) < length {
			return variables.JsonMarshal(Movies), http.StatusOK
		}else {
			return variables.JsonMarshal(Movies[:length]), http.StatusOK
		}
	}else {
		startIndex := 30+(index*10)
		if startIndex > len(Movies) {
			return []byte(`{}`), http.StatusNoContent
		}else if length > len(Movies) {
			return variables.JsonMarshal(Movies[startIndex:]), http.StatusOK
		}else {
			return variables.JsonMarshal(Movies[startIndex:length]), http.StatusOK
		}
	}
}