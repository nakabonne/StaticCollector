package models

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session
var mongoDB *mgo.Database

func dialMongo() (err error) {
	mongoSession, err = mgo.Dial("mongodb://localhost/")
	return
}
func setMongoDB() {
	mongoDB = mongoSession.DB("web_crawler")
}
func closeMongo() {
	mongoSession.Clone()
	log.Println("mongoDB接続がCloseしました。")
}

func getCollection(name string) *mgo.Collection {
	return mongoDB.C(name)
}
