package ratings

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (Rate *Rate) AddToDB() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Ratings")

	if Rate.ExistsByID() {
		Rate.LocalUpdate()
		return
	}

	_, err := collection.InsertOne(ctx, Rate)
	variables.HandleError(err, "ratings", "Rate.AddToDB", "error while inserting rate to the local db")
}


func (rate *Rate) ExistsByID() bool {
	var Rate Rate
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Ratings")

	err := collection.FindOne(ctx, bson.M{"_id": rate.ID}).Decode(&Rate)
	return err == nil
}


func (Rate *Rate) LocalUpdate() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Ratings")

	filter := bson.M{
		"_id": bson.M{
			"$eq": Rate.ID, // check if bool field has value of 'false'
		},
	}

	update := bson.M{"$set": Rate}

	_, err := collection.UpdateOne(ctx, filter, update)
	variables.HandleError(err, "ratings", "Rate.LocalUpdate", "error while updating rate to the local db")
}