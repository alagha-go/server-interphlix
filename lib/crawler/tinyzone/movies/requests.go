package movies

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func UnmarshalLinkResponse(data []byte) (LinkResponse, error) {
	var r LinkResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LinkResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LinkResponse struct {
	Type    string        `json:"type"`   
	Link    string        `json:"link"`   
	Sources []interface{} `json:"sources"`
	Tracks  []interface{} `json:"tracks"` 
	Title   string        `json:"title"`  
}


func PostRequest(url string, data []byte, header bool, headers ...http.Header) ([]byte, http.Header, error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return []byte(""), nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), nil, err
	}

	if !header {
		return body, nil, nil
	}
	return body, res.Header, nil
}

func GetRequest(url string, header bool, headers ...http.Header) ([]byte, http.Header, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte(""), nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), nil, err
	}

	if !header {
		return body, nil, nil
	}
	return body, res.Header, nil
}