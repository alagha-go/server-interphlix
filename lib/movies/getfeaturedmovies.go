package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)



var (
	Titles []string
)


func CollectMovies() {
	SetPopularMovies()
	SetPopularTvShows()
	CollectFeatured()
}


func GetFeaturedMovies(seed int64) ([]Movie, error) {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"featured": true})
	if err != nil {
		variables.HandleError(err, "movies", "GetPopularMovies", "error while getting data from the database")
	}
	cursor.All(ctx, &Movies)
	Movies = RandomMovies(seed, Movies)
	if len(Movies) > 5 {
		return Movies[:5], nil
	}
	return Movies, nil
}