package drive

import (
	"context"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)


func CreateFile(account accounts.Account, filename string) ([]byte, int) {
	ctx := context.Background()
	secretBody, err := ioutil.ReadFile("./secret1.json")
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), 500
	}
	config, err := google.ConfigFromJSON(secretBody, drive.DriveFileScope)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), 500
	}
	client := config.Client(context.Background(), account.Token)

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), 500
	}
	file := &drive.File{Name: filename}
	file, err = driveService.Files.Create(file).Do()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), 500
	}
	return variables.JsonMarshal(file), 200
}