package movies

import (
	"encoding/json"
	"fmt"
	"interphlix/lib/types"
)


func GetMovieUrls(Movie *types.Movie) {
	for _, Server := range Movie.Servers {
				if Server.Name == "Streamlare" {
					var videodata types.VideoData
					postData := fmt.Sprintf(`{"id": "%s"}`, Server.ID)
					data, _, _ := PostRequest("https://streamlare.com/api/video/stream/get", []byte(postData), false)
					json.Unmarshal(data, &videodata)
					Movie.Urls = append(Movie.Urls, videodata.Result.File)
					Movie.Server = &Server
					Movie.Available = true
				}
			}
}