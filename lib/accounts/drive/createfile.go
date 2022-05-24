package drive

import (
	"context"
	"encoding/json"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)


func CreateFile(account accounts.Account, filename string) ([]byte, int) {
	fileExists := FileExist(account, filename)
	if fileExists {
		return variables.JsonMarshal(variables.Error{Error: "file already exists"}), http.StatusConflict
	}
	ctx := context.Background()
	secretBody, err := ioutil.ReadFile("./secret1.json")
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	config, err := google.ConfigFromJSON(secretBody, drive.DriveFileScope)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	client := config.Client(context.Background(), account.Token)

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	file := &drive.File{Name: filename}
	file, err = driveService.Files.Create(file).Do()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	return variables.JsonMarshal(file), http.StatusOK
}


func FileExist(account accounts.Account, filename string) bool {
	var fileList drive.FileList
	data, _ := GetFiles(account)
	json.Unmarshal(data, &fileList)
	for _, file := range fileList.Files {
		if filename == file.Name {
			return true
		}
	}
	return false
}