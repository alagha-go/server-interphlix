package movies


func (Movie *Movie) SetServer() {
	if Movie.Type == "Movie" {
		for _, Server := range Movie.Servers {
			if Server.Name == "Streamlare" {
				Movie.Server = &Server
			}
		}
	}else if Movie.Type == "Tv-Show" {
		for Sindex, Season := range Movie.Seasons {
			for Eindex := range Season.Episodes {
				Movie.Seasons[Sindex].Episodes[Eindex].SetServer()
			}
		}
	}
}

func (Episode *Episode) SetServer() {
	for _, Server := range Episode.Servers {
		if Server.Name == "Streamlare" {
				Episode.Server = &Server
		}
	}
}