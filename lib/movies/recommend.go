package movies

import "interphlix/lib/variables"


func GetMovies() ([]byte, int) {
	if len(Movies) < 30 {
		return variables.JsonMarshal(Movies), 200
	}
	return variables.JsonMarshal(Movies[:30]), 200
}