package main

import (
	"fmt"
	producer2 "github.com/emiliano080591/go-cookbook/kafka/async/producer"
	"gopkg.in/Shopify/sarama.v1"
	"net/http"
)

func main()  {
	config:=sarama.NewConfig()
	config.Producer.Return.Successes=true
	config.Producer.Return.Errors=true
	producer,err:=sarama.NewAsyncProducer([]string{"localhost:9092"},config)
	if err != nil {
		panic(err)
	}
	defer producer.AsyncClose()

	go producer2.ProcessResponse(producer)

	c:=producer2.KafkaController{producer}
	http.HandleFunc("/",c.Handler)
	fmt.Println("Listening on port:3333")
	panic(http.ListenAndServe(":3333",nil))
}
