package ratings

import (
	"context"
	"interphlix/lib/variables"
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