package ratings

import (
	"context"
	"interphlix/lib/movies"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)



func UpdateRate(Movie *movies.Movie, stars int) {
	var rating float64
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	
	err := collection.FindOne(ctx, bson.M{"_id": Movie.ID}).Decode(Movie)
	if err != nil {
		return
	}

	rating = Movie.Rating + float64(stars)
	rating = rating / 2
	
	filter := bson.M{
		"_id": bson.M{
			"$eq": Movie.ID, // check if bool field has value of 'false'
		},
	}

	update := bson.M{"$set": bson.M{
		"rating": rating,
	}}

	collection.UpdateOne(ctx, filter, update)
}


func ChangeRating(Movie *movies.Movie) {
	var Ratings []Rate
	var rating float64
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	collection1 := variables.Client.Database("Interphlix").Collection("Ratings")

	err := collection1.FindOne(ctx, bson.M{"movie_id":Movie.ID}).Decode(&Ratings)
	if err != nil {
		variables.HandleError(err, "movies", "Movie.ChangeRating", "error while getting ratings from the database")
		return
	}

	for _, rate := range Ratings {
		rating = float64(rate.Stars) + rating
	}

	rating = rating / float64(len(Ratings))

	filter := bson.M{
		"_id": bson.M{
			"$eq": Movie.ID, // check if bool field has value of 'false'
		},
	}

	update := bson.M{"$set": bson.M{
		"rating": rating,
	}}

	collection.UpdateOne(ctx, filter, update)
}