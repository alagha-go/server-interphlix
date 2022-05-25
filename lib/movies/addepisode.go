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


func AddEpisode(MovieID, SeasonID primitive.ObjectID, Episode Episode) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	if strings.Contains(Episode.ID.Hex(), "00000000") {
		Episode.ID = primitive.NewObjectID()
	}
	Movie := FindMovie(MovieID)
	if !Movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: fmt.Sprintf("Movie with ID %s does not exist", MovieID.Hex())}), http.StatusNotFound
	}


	for index := range Movie.Seasons {
		if Movie.Seasons[index].ID == SeasonID {
			for _, episode := range Movie.Seasons[index].Episodes {
				if episode.Code == Episode.Code {
					return variables.JsonMarshal(variables.Error{Error: "episode already exists"}), http.StatusConflict
				}
			}
			Movie.Seasons[index].Episodes = append(Movie.Seasons[index].Episodes, Episode)
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
				variables.HandleError(err, "movies", "AddSeason", "error while adding episode to a tvshow")
				return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
			}
			return []byte(`{"success": true}`), http.StatusOK
		}
	}
	return variables.JsonMarshal(variables.Error{Error: fmt.Sprintf("no season with ID %s", SeasonID.Hex())}), http.StatusNotFound
}