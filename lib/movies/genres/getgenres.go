package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllGenres() []Genre {
	var Genres []Genre
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Genres")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "genres", "GetAllGenres", "error while getting genres from the local db")
	err = cursor.All(ctx, &Genres)
	variables.HandleError(err, "genres", "GetAllGenres", "error while decoding cursor")
	return Genres
}


func GetGenresByType(Type string) []Genre {
	var Genres []Genre
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Genres")

	cursor, err := collection.Find(ctx, bson.M{"types": bson.A{Type}})
	variables.HandleError(err, "genres", "GetAllGenresByType", "error while gettoing  genres from the local db")
	err = cursor.All(ctx, &Genres)
	variables.HandleError(err, "genres", "GetAllGenres", "error while decoding cursor")
	return Genres
}