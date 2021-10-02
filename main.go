package main

import (
	"github.com/emiliano080591/go-cookbook/csv"
	"github.com/emiliano080591/go-cookbook/nulls"
	"log"

	"github.com/emiliano080591/go-cookbook/filedirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		log.Println("the file already exists")
	}

	if err := csv.WriteCSVOutput();err!=nil{
		log.Println(err)
	}

	if err := nulls.BaseEncoding();err!=nil{
		log.Println(err)
	}

	if err := nulls.NullEncoding();err!=nil{
		log.Println(err)
	}
}
