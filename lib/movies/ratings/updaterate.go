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

	if Movie.Raters != 0 {
		rating = rating / 2
	}

	Movie.Raters=+1
	
	filter := bson.M{
		"_id": bson.M{
			"$eq": Movie.ID, // check if bool field has value of 'false'
		},
	}

	update := bson.M{"$set": bson.M{
		"rating": rating,
		"raters": Movie.Raters,
	}}

	collection.UpdateOne(ctx, filter, update)
}


func ChangeRating(Movie *movies.Movie) {
	var Ratings []Rate
	var rating float64
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	collection1 := variables.Client.Database("Interphlix").Collection("Ratings")

	cursor, err := collection1.Find(ctx, bson.M{"movie_id":Movie.ID})
	if err != nil {
		variables.HandleError(err, "movies", "Movie.ChangeRating", "error while getting ratings from the database")
		return
	}

	err = cursor.All(ctx, &Ratings)
	if err != nil {
		variables.HandleError(err, "movies", "Movie.ChangeRating", "error while decoding cursor")
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