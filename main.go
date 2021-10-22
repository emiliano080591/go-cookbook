package main

import (
	"context"
	"fmt"
	"github.com/emiliano080591/go-cookbook/state"
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
	//database_mongo.Exec()
	//MONGO con interface intermedia
	//database_mongo.Exec2()

	//urls:=[]string{
	//	"https://www.google.com",
	//	"https://golang.org",
	//	"https://www.github.com",
	//}
	//c:=async.NewClient(http.DefaultClient,len(urls))
	//
	//async.FetchAll(urls,c)
	//for i:=0;i<len(urls);i++ {
	//	select {
	//	case resp:=<-c.Resp:
	//		fmt.Printf("Status received for:%s:%d\n",resp.Request.URL,resp.StatusCode)
	//	case err:=<-c.Err:
	//		fmt.Printf("Error received:%s\n",err)
	//	}
	//}

	//context.Exec()

	// Channels
	test := []struct {
		operation state.Op
		val1      int64
		val2      int64
	}{
		{state.Add, 3, 4},
		{state.Subtract, 5, 2},
		{state.Multiply, 5, 3},
		{state.Divide, 6, 2},
		{state.Divide, 6, 0},
	}

	in := make(chan *state.WorkRequest,5)
	out := make(chan *state.WorkResponse,5)
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processor(ctx, in, out)
	for _, i := range test {
		req := state.WorkRequest{i.operation, i.val1, i.val2}
		in <- &req
	}

	for i := 0; i < 5; i++ {
		resp := <-out
		fmt.Printf("Request: %v; Result: %v, Error: %vn",
			resp.Wr, resp.Result, resp.Err)
	}
}
