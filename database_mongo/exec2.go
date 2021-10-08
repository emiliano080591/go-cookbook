package database_mongo

import (
	"context"
	"log"
)

func Exec2() error {
	m,err:=NewMongoStorage("localhost","goocookbook","items")
	if err != nil {
		return err
	}
	if err:=PerformOperations(m);err!=nil {
		return err
	}

	if err:=m.Session.DB(m.DB).C(m.Collection).DropCollection();err!=nil {
		return err
	}
	return nil
}

func PerformOperations(m *MongoStorage) error {
	ctx:=context.Background()
	i:=Item{
		Name:  "candles",
		Price: 100,
	}
	if err:=m.Put(ctx,&i);err!=nil {
		return err
	}

	candles,err:=m.GetByName(ctx,"candles");
	if err!=nil {
		return err
	}
	log.Printf("Result: %v",candles)
	return nil
}
