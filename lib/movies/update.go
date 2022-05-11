package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// this function adds a server to a movie if it does not exist
func AddServer(ID primitive.ObjectID, Server Server) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	movie := FindMovie(ID)
	if !movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: "Movie does not exist"}), http.StatusNotFound
	}

	for _, server := range movie.Servers {
		if server.Name == Server.Name && server.ID == Server.ID {
			return variables.JsonMarshal(variables.Error{Error: "server already exists"}), http.StatusConflict
		}
	}

	movie.Servers = append(movie.Servers, Server)

	filter := bson.M{
        "_id": bson.M{
            "$eq": ID, // check if bool field has value of 'false'
        },
    }
	update := bson.M{"$set": bson.M{
		"servers": movie.Servers,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "SetServer", "error while updating movie server")
		return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
	}
	index, err := movie.GetIndex()
	if err != nil {
		collection.FindOne(ctx, bson.M{"_id": movie.ID}).Decode(movie)
		Movies = append(Movies, movie)
	}else {
		Movies[index].Servers = movie.Servers
	}
	return variables.JsonMarshal("success"), http.StatusOK
}

/// this function adds url to movies.Url
func AddUrl(ID primitive.ObjectID, url string) ([]byte, int){
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	movie := FindMovie(ID)
	if !movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: "Movie does not exist"}), http.StatusNotFound
	}

	movie.Urls = append(movie.Urls, url)

	filter := bson.M{
        "_id": bson.M{
            "$eq": ID, // check if bool field has value of 'false'
        },
    }
	update := bson.M{"$set": bson.M{
		"urls": movie.Urls,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "SetServer", "error while updating movie urls")
		return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
	}
	index, err := movie.GetIndex()
	if err != nil {
		collection.FindOne(ctx, bson.M{"_id": movie.ID}).Decode(movie)
		Movies = append(Movies, movie)
	}else {
		Movies[index].Urls = movie.Urls
	}
	return variables.JsonMarshal("success"), http.StatusOK
}