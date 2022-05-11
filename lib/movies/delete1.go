package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteServer(MovieID primitive.ObjectID, SeasonCode, EpisodeCode string, server Server) ([]byte, int) {
	Movie := FindMovie(MovieID)

	if Movie.Type == "Movie" {
		exist := false
		for index, Server := range Movie.Servers {
			if Server.Name == server.Name && Server.Url == server.Url {
				exist = true
				Movie.Servers[index] = Movie.Servers[len(Movie.Servers)-1]
				Movie.Servers = Movie.Servers[:len(Movie.Servers)-1]
			}
		}
		if !exist {
			return variables.JsonMarshal(variables.Error{Error: "server does not exist"}), 404
		}else {
			return Movie.UpdateServers()
		}
	}else {
		exist := false
		for Sindex, Season := range Movie.Seasons {
			if Season.Code == SeasonCode {
				for Eindex, Episode := range Season.Episodes {
					for index, Server := range Episode.Servers {
						if Server.Name == server.Name && Server.Url == server.Url {
							exist = true
							Movie.Seasons[Sindex].Episodes[Eindex].Servers[index] = Movie.Seasons[Sindex].Episodes[Eindex].Servers[len(Movie.Seasons[Sindex].Episodes[Eindex].Servers)-1]
							Movie.Seasons[Sindex].Episodes[Eindex].Servers = Movie.Seasons[Sindex].Episodes[Eindex].Servers[:len(Movie.Seasons[Sindex].Episodes[Eindex].Servers)-1]
						}
					}
				}
			}
		}
		if !exist {
			return variables.JsonMarshal(variables.Error{Error: "server does not exist"}), 404
		}else {
			return Movie.UpdateServers()
		}
	}
}

func (Movie *Movie) UpdateServers() ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	filter := bson.M{
		"_id": bson.M{
			"$eq": Movie.ID, // check if bool field has value of 'false'
		},
	}
	update := bson.M{"$set": bson.M{
		"seasons": Movie.Seasons,
		"servers": Movie.Servers,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "UpdateServers", "error while deleting server")
		return variables.JsonMarshal(variables.Error{Error: "could not delete the server"}), http.StatusInternalServerError
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