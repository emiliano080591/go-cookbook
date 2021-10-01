package main

import (
	"github.com/emiliano080591/go-cookbook/csv"
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
}
