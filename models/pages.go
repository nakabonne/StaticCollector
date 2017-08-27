package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Pages ページを管理
type Pages struct {
	ID        bson.ObjectId `bson:"_id"`
	WordID    int           `bson:"word_id"`
	Title     string        `bson:"title"`
	URL       string        `bson:"url"`
	HTML      string        `bson:"html"`
	Rank      int           `bson:"rank"`
	TargetDay time.Time     `bson:"target_day"`
}

// Insert インサート
func (p *Pages) Insert(session *mgo.Session) {
	db := session.DB("web_crawler")
	col := db.C("pages")
	col.Insert(p)
}

// 検索方法はこちら↓
// http://qiita.com/enokidoK/items/a3aff4c05e494b004ef8

//p := new(models.Pages)
//query := db.C("pages").Find(bson.M{})
//query.One(&p)
