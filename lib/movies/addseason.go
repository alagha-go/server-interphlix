package movies

import (
	"context"
	"fmt"
	"interphlix/lib/variables"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddSeason(MovieID primitive.ObjectID, Season Season) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	if strings.Contains(Season.ID.Hex(), "00000000") {
		Season.ID = primitive.NewObjectID()
	}
	Movie := FindMovie(MovieID)
	if !Movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: fmt.Sprintf("Movie with ID %s does not exist", MovieID.Hex())}), http.StatusNotFound
	}
	for index := range Movie.Seasons {
		if Season.Code == Movie.Seasons[index].Code {
			return variables.JsonMarshal(variables.Error{Error: "Season already exists"}), http.StatusConflict
		}
	}
	Movie.Seasons = append(Movie.Seasons, Season)

	filter := bson.M{
		"_id": bson.M{
			"$eq": MovieID, // check if bool field has value of 'false'
		},
	}

	update := bson.M{"$set": bson.M{
			"seasons": Movie.Seasons,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "AddSeason", "error while adding season to a tvshow")
		return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
	}
	index, err := Movie.GetIndex()
	if err != nil {
		Movies = append(Movies, Movie)
	}else {
		Movies[index].Seasons = Movie.Seasons
	}
	return []byte(`{"success": true}`), http.StatusOK
}