package models

func OpenDB() {
	openMysql()
	dialMongo()
	setMongoDB()
}

func CloseDB() {
	closeMysql()
	closeMongo()
}
