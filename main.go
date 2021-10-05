package main

import (
	"github.com/emiliano080591/go-cookbook/cmd"
	"github.com/emiliano080591/go-cookbook/csv"
	"github.com/emiliano080591/go-cookbook/database"
	"github.com/emiliano080591/go-cookbook/nulls"
	"log"

	"github.com/emiliano080591/go-cookbook/filedirs"
)

func main() {
	//files
	if err := filedirs.Operate(); err != nil {
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
	db,err:=database.Setup()
	if err != nil {
		panic(err)
	}
	if err:=database.Exec(db);err != nil {
		panic(err)
	}
}
