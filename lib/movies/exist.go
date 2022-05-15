package movies

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// check if a movie is in the Movies list
func (movie *Movie) Exists() bool {
	for _, Movie := range Movies {
		if movie.Code == Movie.Code && movie.Title == Movie.Title {
			return true
		}
	}
	return false
}



/// find a specific movie from the Movies list
func FindMovie(ID primitive.ObjectID) Movie {
	for _, Movie := range Movies {
		if Movie.ID == ID {
			return Movie
		}
	}
	return Movie{}
}

func (Movie *Movie) Valid() bool {
	return Movie.Title != ""
}