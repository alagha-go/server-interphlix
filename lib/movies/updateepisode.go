package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddEpisodeServer(MovieID primitive.ObjectID, SeasonCode, EpisodeCode string, Server Server) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	movie := FindMovie(MovieID)
	if !movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: "Movie does not exist"}), http.StatusNotFound
	}

	for sindex, Season := range movie.Seasons {
		if Season.Code == SeasonCode {
			for eindex := range Season.Episodes {
				movie.Seasons[sindex].Episodes[eindex].Servers = append(movie.Seasons[sindex].Episodes[eindex].Servers, Server)
			}
		}
	}

	filter := bson.M{
		"_id": bson.M{
			"$eq": movie.ID, // check if bool field has value of 'false'
		},
	}
	update := bson.M{"$set": bson.M{
		"seasons": movie.Seasons,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "SetServer", "error while updating episode server")
		return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
	}
	index, err := movie.GetIndex()
	if err != nil {
		collection.FindOne(ctx, bson.M{"_id": movie.ID}).Decode(movie)
		Movies = append(Movies, movie)
	}else {
		Movies[index].Seasons = movie.Seasons
	}
	return variables.JsonMarshal(`{"success": true}`), http.StatusOK
}