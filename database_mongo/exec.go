package database_mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"population"`
}

func Exec() error {
	db,err:=Setup()
	if err != nil {
		return err
	}

	conn:=db.DB("gocookbook").C("example")

	if errCon:=conn.Insert(&State{Name: "Washington",Population: 1000},&State{Name: "Texas",Population: 50000});errCon != nil {
		return err
	}

	var s State

	if err:=conn.Find(bson.M{"name":"Texas"}).One(&s);err!=nil {
		return err
	}

	if err:=conn.DropCollection();err!=nil{
		return err
	}

	fmt.Printf("State : %v",s)
	return nil
}
