package history

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type History struct {
	ID							primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	AccountID					primitive.ObjectID							`json:"account_id,omitempty" bson:"account_id,omitempty"`
	MovieID						primitive.ObjectID							`json:"movie_id,omitempty" bson:"movie_id,omitempty"`
	Percentage					float64										`json:"percentage,omitempty" bson:"percentage,omitempty"`
	Episodes					[]Episode									`json:"episodes,omitempty" bson:"episodes,omitempty"`
	Episode						*Episode									`json:"episode,omitempty" bson:"episode,omitempty"`
	LastTimeWatched				time.Time									`json:"last_time_watched,omitempty" bson:"last_time_watched,omitempty"`
}

type Episode struct {
	ID							primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	SeasonIndex					int											`json:"season_index,omitempty" bson:"season_index,omitempty"`
	Percentage					float64										`json:"percentage,omitempty" bson:"percentage,omitempty"`
	LastTimeWatched				time.Time									`json:"last_time_watched,omitempty" bson:"last_time_watched,omitempty"`
}