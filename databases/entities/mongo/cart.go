package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Cart struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Content      string        `bson:"content"`
	Title        string        `bson:"title"`
	CreatedBy    string        `bson:"created_by"`
	CreatedDate  time.Time     `bson:"created_date"`
	ModifiedBy   string        `bson:"modified_by"`
	ModifiedDate time.Time     `bson:"modified_date"`
}
