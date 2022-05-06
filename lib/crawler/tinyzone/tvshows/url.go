package tvshows

import (
	"encoding/json"
	"fmt"
	"interphlix/lib/types"
)


func GetTvShowUrls(TvShow *types.Movie) {
	for sindex, Season := range TvShow.Seasons {
		for eindex, Episode := range Season.Episodes {
			for _, Server := range Episode.Servers {
				if Server.Name == "Streamlare" {
					var videodata types.VideoData
					postData := fmt.Sprintf(`{"id": "%s"}`, Server.ID)
					data, _, _ := PostRequest("https://streamlare.com/api/video/stream/get", []byte(postData), false)
					json.Unmarshal(data, &videodata)
					Episode.Server = &Server
					TvShow.Seasons[sindex].Episodes[eindex].Urls = append(Episode.Urls, videodata.Result.File)
				}
			}
		}
	}
}