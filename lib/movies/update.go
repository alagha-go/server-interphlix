package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SetServer(ID primitive.ObjectID, Server Server) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	filter := bson.M{
        "_id": bson.M{
            "$eq": ID, // check if bool field has value of 'false'
        },
    }
	update := bson.M{
		"server": Server,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "SetServer", "error while updating movie server")
		return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal("success"), http.StatusOK
}


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
	update := bson.M{
		"urls": movie.Urls,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "SetServer", "error while updating movie urls")
		return variables.JsonMarshal(variables.Error{Error: "could not update movie"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal("success"), http.StatusOK
}