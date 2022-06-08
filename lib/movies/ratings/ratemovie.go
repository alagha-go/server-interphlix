package ratings

import (
	"context"
	"interphlix/lib/variables"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (Rate *Rate) RateMovie() (string, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Ratings")
	if Rate.Exists() {
		return string(variables.JsonMarshal(variables.Error{Error: "rate already exists"})), http.StatusConflict
	}
	Rate.ID = primitive.NewObjectID()
	Rate.TimeRated = time.Now()
	_, err := collection.InsertOne(ctx, Rate)
	if err != nil {
		variables.HandleError(err, "ratings", "RateMovie", "error while inserting rate to the database")
		return string(variables.JsonMarshal(variables.Error{Error: "could save your rate"})), http.StatusInternalServerError
	}
	return string(variables.JsonMarshal(Rate)), http.StatusCreated
}


func (rate *Rate) Exists() bool {
	var Rate Rate
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Ratings")

	err := collection.FindOne(ctx, bson.M{"movie_id": rate.MovieID, "account_id": rate.AccountID}).Decode(&Rate)
	return err == nil
}


func (Rate *Rate) Update() (string, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Ratings")

	if !Rate.ExistsByID() {
		return `{"error": "rating does not exist"}`, http.StatusNotFound
	}
	
	filter := bson.M{
		"_id": bson.M{
			"$eq": Rate.ID, // check if bool field has value of 'false'
		},
	}
	Rate.TimeRated = time.Now()

	update := bson.M{"$set": bson.M{
		"stars": Rate.Stars,
		"review": Rate.Review,
		"time_rated": Rate.TimeRated,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "ratings", "Rate.Update", "error while updating a rating")
		return `{"error": "internal server error"}`, http.StatusInternalServerError
	}
	return string(variables.JsonMarshal(Rate)), http.StatusOK
}