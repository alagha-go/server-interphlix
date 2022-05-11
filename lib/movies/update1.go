package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SetServer(MovieID primitive.ObjectID, Servername, ServerID string) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	movie := FindMovie(MovieID)
	if !movie.Valid() {
		return variables.JsonMarshal(variables.Error{Error: "Movie does not exist"}), http.StatusNotFound
	}

	for _, server := range movie.Servers {
		if server.Name == Servername && server.ID == ServerID {
			filter := bson.M{
				"_id": bson.M{
					"$eq": MovieID, // check if bool field has value of 'false'
				},
			}
			update := bson.M{"$set": bson.M{
				"server": server,
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
				Movies[index].Server = movie.Server
			}
			return variables.JsonMarshal("success"), http.StatusOK
		}
	}
	return variables.JsonMarshal(variables.Error{Error: "Sevrer does not exist"}), http.StatusNotFound
}