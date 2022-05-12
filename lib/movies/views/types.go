package views

import "go.mongodb.org/mongo-driver/bson/primitive"


type View struct {
	ID									primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	MovieID								primitive.ObjectID								`json:"movie-id,omitempty" bson:"movie-id,omitempty"`
	AccountID							primitive.ObjectID								`json:"account-id,omitempty" bson:"account-id,omitempty"`
	UserID								primitive.ObjectID								`json:"user-id,omitempty" bson:"user-id,omitempty"`
	PercentageViewed					float64											`json:"percentage-viewed,omitempty" bson:"percentage-viewed,omitempty"`
}