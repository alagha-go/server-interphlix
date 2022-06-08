package movies

import (
	"context"
	"interphlix/lib/movies/casts"
	"interphlix/lib/movies/genres"
	"interphlix/lib/movies/ratings"
	"interphlix/lib/movies/types"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)




func Main() {
	LoadMovies()
	types.Main()
	genres.Main()
	casts.Main()
	ratings.Main()
	Listener()
}


func LoadMovies() {
	var documents []interface{}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	collection1 := variables.Client1.Database("Interphlix").Collection("Movies")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "movies", "LoadMoviesFromDB", "error while getting movies from the remote database")
		return
	}
	cursor.All(ctx, &documents)
	collection1.Drop(ctx)
	_, err = collection1.InsertMany(ctx, documents)
	variables.HandleError(err, "movies", "LoadMoviesFromDB", "error while inserting movies to the local database")
}