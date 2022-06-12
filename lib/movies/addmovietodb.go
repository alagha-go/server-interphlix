package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (movie *Movie) AddToDB() {
	var dbmovie Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(ctx, bson.M{"_id": movie.ID}).Decode(&dbmovie)
	if err != nil {
		movie.AddIndex()
		collection.InsertOne(ctx, movie)
		return
	}
	filter := bson.M{
		"_id": bson.M{
			"$eq": movie.ID, // check if bool field has value of 'false'
		},
	}
	update := bson.M{"$set": movie}
	collection.UpdateOne(ctx, filter, update)
}