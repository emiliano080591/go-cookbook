package database_mongo

import "gopkg.in/mgo.v2"

func Setup() (*mgo.Session,error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session,nil
}
