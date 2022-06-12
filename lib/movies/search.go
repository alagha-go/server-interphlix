package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"github.com/blevesearch/bleve/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SearchMovies(querry, Type, genre string, round int) ([]byte, int) {
	var length int = 50
	if round != 0 {
		length = (round*20) + 50
	}
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	query := bleve.NewQueryStringQuery(querry)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := Index.Search(searchRequest)
	if err != nil {
		variables.HandleError(err, "movies", "SearchMovies", "error while searching data")
		return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
	}
	for _, Hit := range searchResult.Hits {
		var Movie Movie
		ID, _ := primitive.ObjectIDFromHex(Hit.ID)
		err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Movie)
		if err == nil {
			if Type != "" {
				if Movie.Type == Type {
					for _, Genre := range Movie.Genres {
						if Genre == genre {
							Movies = append(Movies, Movie)
							break
						}
					}
				}
			}else {
				Movies = append(Movies, Movie)
			}
		}
	}

	if len(Movies) > length {
		return variables.JsonMarshal(Movies), http.StatusOK
	}

	return variables.JsonMarshal(Movies[:length]), http.StatusOK
}