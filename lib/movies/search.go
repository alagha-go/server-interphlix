package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func SearchMovies(querry, Type, genre string, round int) ([]byte, int) {
	start := 0
	var length int = 50
	if round != 0 {
		if round == 1{
			start = 50
		}else {
			start = ((round-1)*20) + 50
		}
		length = (round*20) + 50
	}
	var Movies []Movie
	var NewMovies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	projection := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	opts := options.Find().SetSort(sort).SetProjection(projection)

	if Type != "" {
		var innerMovies []Movie
		cursor, err := collection.Find(ctx, bson.M{"type": Type, "$text": bson.M{"$search": querry}}, opts)
		if err != nil {
			return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
		}
		cursor.All(ctx, &innerMovies)
		if genre != "" {
			for _, Movie := range innerMovies {
				for _, Genre := range Movie.Genres {
					if Genre == genre {
						Movies = append(Movies, Movie)
					}
				}
			}
		}
	}else {
		cursor, err := collection.Find(ctx, bson.M{"$text": bson.M{"$search": querry}}, opts)
		if err != nil {
			return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
		}
		cursor.All(ctx, &Movies)
	}

	for index, movie := range Movies {
		if index >= length {
			return variables.JsonMarshal(NewMovies), http.StatusOK
		}else if index >= start && index < length{
			NewMovie := Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl}
			NewMovies = append(NewMovies, NewMovie)
		}
	}

	return variables.JsonMarshal(NewMovies), http.StatusOK
}