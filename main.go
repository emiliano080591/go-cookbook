package main

import (
	"github.com/emiliano080591/go-cookbook/database_redis"
	"log"
)

func main() {
	//files
	/*if err := filedirs.Operate(); err != nil {
		log.Println("the file already exists")
	}
	//CSV
	if err := csv.WriteCSVOutput();err!=nil{
		log.Println(err)
	}

	if err := nulls.BaseEncoding();err!=nil{
		log.Println(err)
	}
	//Parse with sql
	if err := nulls.NullEncoding();err!=nil{
		log.Println(err)
	}
	//flags
	c:=cmd.Config{}
	c.Setup()
	log.Println(c)

	//tags.EmptyStruct()
	//tags.FullStruct()
	//db,err:= database_mysql.Setup()
	//if err != nil {
	//	panic(err)
	//}
	//if err:= database_mysql.Exec(db);err != nil {
	//	panic(err)
	//}
*/
	client,err:=database_redis.Setup()
	log.Println(client.Ping().String(),err)
}
