package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func LoadGenres() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "movies", "LoadGenres", "error while getting genres from the Database")
	err = cursor.All(ctx, &Genres)
	variables.HandleError(err, "movies", "LoadGenres", "error while decoding genres from the cursor")
}


func (Genre *Genre) UpdateGenre() {
	for index, genre := range Genres {
		if genre.ID == Genre.ID {
			for _, Type := range genre.Types {
				if Type == Genre.Type {
					return
				}
			}
			Genres[index].Types = append(Genres[index].Types, Genre.Type)
		}
	}
}



func (Genre *Genre) AddGenre() {
	Genres = append(Genres, *Genre)
}