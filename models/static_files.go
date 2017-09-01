package models

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Pages ページを管理
type StaticFiles struct {
	ID        bson.ObjectId `bson:"_id"`
	WordID    int           `bson:"word_id"`
	PageID    int           `bson:"page_id"`
	Title     string        `bson:"title"`
	HTML      string        `bson:"html"`
	Rank      int           `bson:"rank"`
	TargetDay time.Time     `bson:"target_day"`
}

func getCollection(session *mgo.Session) *mgo.Collection {
	db := session.DB("web_crawler")
	col := db.C("static_files")
	return col
}

// Insert インサート
func (p *StaticFiles) Insert(session *mgo.Session) {
	col := getCollection(session)
	col.Insert(p)
}

// TODO Find系はinterfaceとか使って一元化する
func FindStaticFilesByPageWord(pageID int, wordID int, session *mgo.Session) []*StaticFiles {
	staticFiles := make([]*StaticFiles, 0)
	col := getCollection(session)
	if err := col.Find(bson.M{
		"page_id": pageID,
		"word_id": wordID,
	}).All(&staticFiles); err != nil {
		log.Fatal("エラー", err)
	}
	return staticFiles
}

func FindStaticFilesByPageWordTargetday(pageID int, wordID int, targetDay time.Time, session *mgo.Session) *StaticFiles {
	staticFiles := make([]*StaticFiles, 0)
	col := getCollection(session)
	if err := col.Find(bson.M{
		"page_id":    pageID,
		"word_id":    wordID,
		"target_day": targetDay,
	}).All(&staticFiles); err != nil {
		log.Fatal("エラー", err)
	}
	return staticFiles[0]
}

// 検索方法はこちら↓
// http://qiita.com/enokidoK/items/a3aff4c05e494b004ef8

//p := new(models.Pages)
//query := db.C("pages").Find(bson.M{})
//query.One(&p)
