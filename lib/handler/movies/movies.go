package movies

import (
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMoviesByTypeAndGenre(res http.ResponseWriter, req *http.Request) {
	var Movies []movies.Movie
	res.Header().Set("content-type", "application/json")
	Type := mux.Vars(req)["type"]
	Genre := mux.Vars(req)["genre"]

	for _, Movie := range movies.Movies {
		if Movie.Type == Type {
			for _, genre := range Movie.Genres {
				if genre == Genre {
					Movies = append(Movies, Movie)
				}
			}
		}
	}

	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(Movies))
}