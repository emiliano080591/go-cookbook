package main

import (
	"fmt"
	"github.com/emiliano080591/go-cookbook/async"
	"github.com/emiliano080591/go-cookbook/database_mongo"
	"net/http"
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
	//REDIS
	//client,err:=database_redis.Setup()
	//if err:=database_redis.Exec();err!=nil{
	//	panic(err)
	//}
	//if err:=database_redis.Sort();err!=nil{
	//	panic(err)
	//}

	//MONGO
	database_mongo.Exec()
	//MONGO con interface intermedia
	database_mongo.Exec2()

	urls:=[]string{
		"https://www.google.com",
		"https://golang.org",
		"https://www.github.com",
	}
	c:=async.NewClient(http.DefaultClient,len(urls))

	async.FetchAll(urls,c)
	for i:=0;i<len(urls);i++ {
		select {
		case resp:=<-c.Resp:
			fmt.Printf("Status received for:%s:%d\n",resp.Request.URL,resp.StatusCode)
		case err:=<-c.Err:
			fmt.Printf("Error received:%s\n",err)
		}
	}
}
