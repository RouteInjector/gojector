package mongo

import (
	"gopkg.in/mgo.v2"
	"github.com/RouteInjector/gojector/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/RouteInjector/gojector/infrastructure/conf"
	"fmt"
	"strings"
)

type MongoHandler struct {
	Session *mgo.Session  // MongoDB Session
	Db      *mgo.Database // MongoDB Database
}

func NewMongoHandler(conf *conf.Database) *MongoHandler {
	mHandler := &MongoHandler{}
	mgo.SetDebug(true)
	session, err := mgo.Dial(conf.Endpoint)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
	}
	mHandler.Session = session
	mHandler.Db = session.DB(conf.Name)
	return mHandler
}

func (m *MongoHandler) StopDatabase() {
	m.Session.Close();
}

type ModelWrapper struct {
	Model      *model.Model
	collection *mgo.Collection
}

func (mhandler MongoHandler) WrapModel(model *model.Model) *ModelWrapper {
	w := &ModelWrapper{
		Model: model,
		collection: mhandler.Db.C(strings.ToLower(model.Name)),
	}
	return w
}

func (w *ModelWrapper) Insert(data interface{}) (err error) {
	return w.collection.Insert(data)
}

func (w *ModelWrapper) FindOne(id interface{}) (doc interface{}, err error) {
	fmt.Println("FindOne -> " + w.Model.Name + " -> " + w.Model.ID + " -> " + id.(string))
	err = w.collection.Find(bson.M{w.Model.ID:id}).One(&doc)
	fmt.Println(doc)
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