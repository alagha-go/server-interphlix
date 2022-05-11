package movies

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
		variables.HandleError(err, "movies", "Movie.Upload", "could not upload genre to the Database")
		return err
	}
	Genre.AddGenre()
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

	if Genre.Afro {
		update = bson.M{"$set": bson.M{
			"afro": true,
		}}
	}else if Genre.Fanproj {
		update = bson.M{"$set": bson.M{
			"fanproj": true,
		}}
	}else if Genre.TvShow {
		update = bson.M{"$set": bson.M{
			"tv-show": true,
		}}
	}else {
		update = bson.M{"$set": bson.M{
			"movie": true,
		}}
	}

	genre := Genre.Find()
	if Genre.Afro == genre.Afro && Genre.TvShow == genre.TvShow && Genre.Movie == genre.Movie && Genre.Fanproj == genre.Fanproj {
		return nil
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "movies", "Movie.Upload", "could not update genre to the Database")
		return err
	}
	Genre.UpdateGenre()
	return nil
}

/// cgheck if given genre is valid
func (Genre *Genre) Valid() bool {
	return Genre.Title != ""
}