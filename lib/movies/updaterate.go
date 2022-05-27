package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)



func (Movie *Movie) UpdateRate(stars int) {
	var rating float64
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	
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