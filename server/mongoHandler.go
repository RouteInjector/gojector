package server

import (
	"gopkg.in/mgo.v2"
)

type MongoHandler struct {
	Session *mgo.Session  // MongoDB Session
	Db      *mgo.Database // MongoDB Database
}

func NewMongoHandler(url, dbName string) *MongoHandler {
	mHandler := &MongoHandler{}
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	mHandler.Session = session
	mHandler.Db = session.DB(dbName)
	return mHandler
}

func (m *MongoHandler) StopDatabase(){
	m.Session.Close();
}