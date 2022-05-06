package types

import "encoding/json"


func UnmarshalVideoData(data []byte) (VideoData, error) {
	var r VideoData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VideoData) Marshal() []byte {
	return JsonMarshal(r)
}

type VideoData struct {
	Status  string `json:"status"` 
	Message string `json:"message"`
	Type    string `json:"type"`   
	Result  Result `json:"result"` 
}

type Result struct {
	File  string `json:"file"` 
	Type  string `json:"type"` 
	Title string `json:"title"`
}