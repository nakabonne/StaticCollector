package models

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Pages ページを管理
type StaticFile struct {
	ID        bson.ObjectId `bson:"_id"`
	WordID    int           `bson:"word_id"`
	PageID    int           `bson:"page_id"`
	Title     string        `bson:"title"`
	HTML      string        `bson:"html"`
	Rank      int           `bson:"rank"`
	TargetDay time.Time     `bson:"target_day"`
}

//func getCollection(session *mgo.Session) *mgo.Collection {
//	db := session.DB("web_crawler")
//	col := db.C("static_files")
//	return col
//}

// Insert インサート
func (p *StaticFile) Insert() {
	col := getCollection("static_files")
	col.Insert(p)
}

// TODO Find系はinterfaceとか使って一元化する
func FindStaticFilesByPageWord(pageID int, wordID int) []*StaticFile {
	staticFiles := make([]*StaticFile, 0)
	col := getCollection("static_files")
	if err := col.Find(bson.M{
		"page_id": pageID,
		"word_id": wordID,
	}).All(&staticFiles); err != nil {
		log.Fatal("エラー", err)
	}
	return staticFiles
}

func FindStaticFilesByPageWordTargetday(pageID int, wordID int, targetDay time.Time) (staticFile *StaticFile, err error) {
	staticFiles := make([]*StaticFile, 0)
	col := getCollection("static_files")
	if err = col.Find(bson.M{
		"page_id":    pageID,
		"word_id":    wordID,
		"target_day": targetDay,
	}).All(&staticFiles); err != nil {
		return
	}
	if len(staticFiles) >= 1 {
		staticFile = staticFiles[0]
	} else {
		staticFile = &StaticFile{}
	}
	return
}
