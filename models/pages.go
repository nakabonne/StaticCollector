package models

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
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

func (p *Pages) Insert() {
	session := session()
	defer session.Clone()
	db := session.DB("web_crawler")
	col := db.C("pages")
	col.Insert(p)
}

func session() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost/")
	if err != nil {
		log.Fatal("エラー", err)
	}
	return session
}
