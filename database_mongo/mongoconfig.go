package database_mongo

import (
	"context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoStorage struct {
	*mgo.Session
	DB         string
	Collection string
}

func (m *MongoStorage) GetByName(ctx context.Context,name string) (*Item,error){
	c:=m.Session.DB(m.DB).C(m.Collection)
	var i Item
	if err:=c.Find(bson.M{"name":name}).One(&i);err != nil {
		return nil, err
	}
	return &i, nil
}

func (m *MongoStorage) Put(ctx context.Context,i *Item) error{
	c:=m.Session.DB(m.DB).C(m.Collection)
	return c.Insert(i)
}

func NewMongoStorage(connection,db,collection string) (*MongoStorage,error) {
	session,err:=mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	ms:=MongoStorage{
		session,
		db,
		collection,
	}
	return &ms,nil
}
