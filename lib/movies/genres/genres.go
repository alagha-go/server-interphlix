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
	for gindex, genre := range Genres {
		if genre.ID == Genre.ID {
			if Genre.Afro {
				Genres[gindex].Afro = true
			}else if Genre.Fanproj {
				Genres[gindex].Fanproj = true
			}else if Genre.TvShow {
				Genres[gindex].TvShow = true
			}else {
				Genres[gindex].Movie = true
			}
		}
	}
}



func (Genre *Genre) AddGenre() {
	Genres = append(Genres, *Genre)
}