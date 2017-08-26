package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Pages ページを管理
type Pages struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string        `bson:"title"`
	URL       string        `bson:"url"`
	HTML      string        `bson:"html"`
	Rank      int           `bson:"rank"`
	TargetDay time.Time     `bson:"target_day"`
}
