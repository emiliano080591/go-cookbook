package main

import (
	"log"

	"github.com/emiliano080591/go-cookbook/filedirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		log.Println("the file already exists")
	}
}
