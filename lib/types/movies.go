package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Movie struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Code										string									`json:"code,omitempty" bson:"code,omitempty"`
	Title										string									`json:"title,omitempty" bson:"title,omitempty"`
    PageUrl										string									`json:"page_url,omitempty" bson:"page_url,omitempty"`
    ImageUrl									string									`json:"image_url,omitempty" bson:"image_url"`
    Released									*time.Time								`json:"released,omitempty" bson:"released"`
    Urls										[]string								`json:"urls,omitempty" bson:"urls,omitempty"`
    Genres										[]string								`json:"genre,omitempty" bson:"genre,omitempty"`
    Servers										[]Server								`json:"servers,omitempty" bson:"servers,omitempty"`
    Casts										[]string								`json:"casts,omitempty" bson:"casts,omitempty"`
    Duration									int										`json:"duration,omitempty" bson:"duration,omitempty"`
    Countries									[]string								`json:"countries,omitempty" bson:"countries,omitempty"`
    Producers									[]string								`json:"producers,omitempty" bson:"producers,omitempty"`
    Description									string									`json:"description,omitempty" bson:"description,omitempty"`
}


type Season struct {
	Name										string									`json:"name,omitempty" bson:"name,omitempty"`
	Episodes									[]Episode								`json:"episodes,omitempty" bson:"episodes,omitempty"`
}


type Episode struct {
	Name										string									`json:"name,omitempty" bson:"name,omitempty"`
	Servers										[]Server								`json:"servers,omitempty" bson:"servers,omitempty"`
}


type Server struct {
    Name                                           string                               `json:"name"`
    ID                                             string                               `json:"id"`
    WatchID                                        string                               `json:"watch_id"`
    Url                                            string                               `json:"url"`
}