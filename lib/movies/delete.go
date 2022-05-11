package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteUrls(MovieID primitive.ObjectID, SeasonCode, EpisodeCode string, urls ...string) ([]byte, int) {
	Movie := FindMovie(MovieID)
	if !Movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: "Movie does not exist"}), 404
	}
	if Movie.Type == "Movie" {
		exist := false
		for index, Url := range Movie.Urls {
			for _, url := range urls {
				if Url == url {
					exist = true
					Movie.Urls[index] = Movie.Urls[len(Movie.Urls)-1]
					Movie.Urls = Movie.Urls[:len(Movie.Urls)-1]
				}
			}
		}
		if !exist {
			return variables.JsonMarshal(variables.Error{Error: "urls do not exist"}), 404
		}else {
			return Movie.RemoveUrl()
		}
	}else {
		exist := false
		for Sindex, Season := range Movie.Seasons {
			if Season.Code == SeasonCode {
				for Eindex, Episode := range Season.Episodes {
					if Episode.Code == EpisodeCode {
						for index, Url := range Episode.Urls {
							for _, url := range urls {
								if Url == url {
									exist = true
									Movie.Seasons[Sindex].Episodes[Eindex].Urls[index] = Movie.Seasons[Sindex].Episodes[Eindex].Urls[len(Movie.Seasons[Sindex].Episodes[Eindex].Urls)-1]
									Movie.Seasons[Sindex].Episodes[Eindex].Urls = Movie.Seasons[Sindex].Episodes[Eindex].Urls[:len(Movie.Seasons[Sindex].Episodes[Eindex].Urls)-1]
								}
							}
						}
					}
				}
			}
		}
		if !exist {
			return variables.JsonMarshal(variables.Error{Error: "urls do not exist"}), 404
		}else {
			return Movie.RemoveUrl()
		}
	}
}


func (Movie *Movie) RemoveUrl() ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	filter := bson.M{
		"_id": bson.M{
			"$eq": Movie.ID, // check if bool field has value of 'false'
		},
	}
	update := bson.M{"$set": bson.M{
		"seasons": Movie.Seasons,
		"urls": Movie.Urls,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "SetServer", "error while deleting movie urls")
		return variables.JsonMarshal(variables.Error{Error: "could not not delete urls"}), http.StatusInternalServerError
	}
	index, err := Movie.GetIndex()
	if err != nil {
		collection.FindOne(ctx, bson.M{"_id": Movie.ID}).Decode(Movie)
		Movies = append(Movies, *Movie)
	}else {
		Movies[index] = *Movie
	}
	return []byte(`{"success": true}`), http.StatusOK
}