package genres

import "go.mongodb.org/mongo-driver/bson/primitive"

type Genre struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Title										string									`json:"title,omitempty" bson:"title,omitempty"`
	TvShow										bool									`json:"tv-show,omitempty" bson:"tv-show,omitempty"`
	Movie										bool									`json:"movie,omitempty" bson:"movie,omitempty"`
	Afro										bool									`json:"afro,omitempty" bson:"afro,omitempty"`
	Fanproj										bool									`json:"fanproj,omitempty" bson:"fanproj,omitempty"`
}