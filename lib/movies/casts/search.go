package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"github.com/blevesearch/bleve/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SearchCasts(querry string, round int) ([]byte, int) {
	var length int = 50
	if round != 0 {
		length = (round*20) + 50
	}
	var Casts []Cast
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	query := bleve.NewQueryStringQuery(querry)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := Index.Search(searchRequest)
	if err != nil {
		variables.HandleError(err, "casts", "SearchCasts", "error while searching data")
		return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
	}

	for _, Hit := range searchResult.Hits {
		var Cast Cast
		ID, _ := primitive.ObjectIDFromHex(Hit.ID)
		err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Cast)
		if err == nil {
			Casts = append(Casts, Cast)
		}
	}

	if len(Casts) < length {
		return variables.JsonMarshal(Casts), http.StatusOK
	}

	return variables.JsonMarshal(Casts[:length]), http.StatusOK
}