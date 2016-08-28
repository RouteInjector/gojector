package mongo

import (
	"gopkg.in/mgo.v2"
	"github.com/RouteInjector/gojector/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/RouteInjector/gojector/infrastructure/conf"
)


type MongoHandler struct {
	Session *mgo.Session  // MongoDB Session
	Db      *mgo.Database // MongoDB Database
}

func NewMongoHandler(conf *conf.Database) *MongoHandler {
	mHandler := &MongoHandler{}
	session, err := mgo.Dial(conf.Endpoint)
	if err != nil {
		panic(err)
	}
	mHandler.Session = session
	mHandler.Db = session.DB(conf.Name)
	return mHandler
}

func (m *MongoHandler) StopDatabase(){
	m.Session.Close();
}

type ModelWrapper struct {
	Model      *model.Model
	collection *mgo.Collection
}

func (mhandler MongoHandler) WrapModel(model *model.Model) *ModelWrapper {
	w := &ModelWrapper{
		Model: model,
		collection: mhandler.Db.C(model.Name),
	}
	return w
}

func (w *ModelWrapper) FindOne(id interface{}) (doc interface{}, err error) {
	err = w.collection.Find(bson.M{w.Model.ID:id}).One(&doc)
	return doc, err
}

func (w *ModelWrapper) Delete(id interface{}) (err error) {
	err = w.collection.Remove(bson.M{w.Model.ID:id})
	return err
}

func (w *ModelWrapper) Search(query bson.M) (docs []interface{}, err error) {
	err = w.collection.Find(query).One(&docs)
	return docs, err
}