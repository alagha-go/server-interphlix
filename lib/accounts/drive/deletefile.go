package drive

import (
	"context"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)


func DeleteFile(account accounts.Account, fileid string) ([]byte, int) {
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

	err = driveService.Files.Delete(fileid).Do()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	return variables.JsonMarshal([]byte(`{"suuccess": true}`)), http.StatusOK
}