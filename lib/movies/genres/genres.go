package genres

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func LoadGenres() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Genres")
	collection := variables.Client1.Database("Interphlix").Collection("Genres")

	cursor, err := collection1.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "genres", "LoadGenres", "error while getting genres from the Database")
		return
	}
	err = cursor.All(ctx, &documents)
	if err != nil {
		variables.HandleError(err, "genres", "LoadGenres", "error while decoding genres from the cursor")
		return
	}
	collection.Drop(ctx)
	collection.InsertMany(ctx, documents)
}