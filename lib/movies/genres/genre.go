package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// check if a genre exists in our inmemory Genres
func (genre *Genre) Exists() bool {
	for _, Genre := range Genres {
		if Genre.Title == genre.Title {
			return true
		}
	}
	return false
}


// finc and return genre with the same title
func (genre *Genre) Find() Genre {
	for _, Genre := range Genres {
		if Genre.Title == genre.Title {
			return Genre
		}
	}
	return Genre{}
}


/// upload a genre to the Database
func (Genre *Genre) Upload() error {
	if Genre.Title == "" || Genre.Title == " " {
		return nil
	}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	_, err := collection.InsertOne(ctx, Genre)
	if err != nil {
		variables.HandleError(err, "genres", "Genre.Upload", "could not upload genre to the Database")
		return err
	}
	return nil
}


/// update genre if it needs update
func (Genre *Genre) Update() error{
	var filter primitive.M
	var update primitive.M
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	filter = bson.M{
		"title": Genre.Title,
	}

	genre := Genre.Find()

	for _, Type := range genre.Types {
		if Type == Genre.Type {
			return nil
		}
	}

	genre.Types = append(genre.Types, Genre.Type)

	update = bson.M{"$set": bson.M{
			"types": genre.Types,
	}}


	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "genres", "Genre.Update", "could not update genre to the Database")
		return err
	}
	return nil
}

/// cgheck if given genre is valid
func (Genre *Genre) Valid() bool {
	return Genre.Title != ""
}