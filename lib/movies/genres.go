package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func LoadGenres() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "LoadGenres", "error while getting genres from the Database")
	err = cursor.All(ctx, &Genres)
	variables.HandleError(err, "LoadGenres", "error while decoding genres from the cursor")
}


func ListenGenres() {
	var genres []Genre
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "ListenMovies", "error while listening for changes in Genres collection")
	err = cursor.All(ctx, &genres)
	variables.HandleError(err, "ListenMovies", "error while decoding Genres from the cursor")
	for _, genre := range genres {
		if genre.Exists() {
			for genreIndex, Genre := range Genres {
				if Genre.ID == genre.ID {
					Genres[genreIndex] = genre
				}
			}
		}else {
			Genres = append(Genres, genre)
		}
	}
	ListenGenres()
}