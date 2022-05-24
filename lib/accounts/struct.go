package accounts

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)


type Account struct {
	ID									primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	GoogleID							string										`json:"id,omitempty" bson:"id,omitempty"`
	Email								string										`json:"email,omitempty" bson:"email,omitempty"`
	EmailVerified						bool										`json:"verified_email,omitempty" bson:"verified_email,omitempty"`
	Name								string										`json:"name,omitempty" bson:"name,omitempty"`
	GivenName							string										`json:"given_name,omitempty" bson:"given_name,omitempty"`
	FamilyName							string										`json:"family_name,omitempty" bson:"family_name,omitempty"`
	Picture								string										`json:"picture,omitempty" bson:"picture,omitempty"`
	Locale								string										`json:"locale,omitempty" bson:"locale,omitempty"`
	Token								*oauth2.Token								`json:"token,omitempty" bson:"token,omitempty"`
	Deleted								bool										`json:"deleted,omitempty" bson:"deleted,omitempty"`
}