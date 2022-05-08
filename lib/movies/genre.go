package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (genre *Genre) Exists() bool {
	var DBGenre Genre
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	err := collection.FindOne(ctx, bson.M{"title": genre.Title}).Decode(&DBGenre)
	if err != nil {
		return false
	}
	return true
}

func (Genre *Genre) Upload() error {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	_, err := collection.InsertOne(ctx, Genre)
	if err != nil {
		variables.HandleError(err, "Movie.Upload", "could not upload genre to the Database")
		return err
	}
	return nil
}

func (Genre *Genre) Update() error{
	var filter primitive.M
	var update primitive.M
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	filter = bson.M{
		"title": Genre.Title,
	}

	if Genre.Afro {
		update = bson.M{
			"afro": true,
		}
	}else if Genre.Fanproj {
		update = bson.M{
			"fanproj": true,
		}
	}else if Genre.TvShow {
		update = bson.M{
			"tv-show": true,
		}
	}else {
		update = bson.M{
			"movie": true,
		}
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "Movie.Upload", "could not update genre to the remoteDB")
		return err
	}
	return nil
}