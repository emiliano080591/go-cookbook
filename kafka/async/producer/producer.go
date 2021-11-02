package producer

import (
	"gopkg.in/Shopify/sarama.v1"
	"log"
)

func ProcessResponse(producer sarama.AsyncProducer)  {
	for {
		select {
		case result:=<-producer.Successes():
			log.Printf("> message: \"%s\" sent to partition %d at offset %d\n", result.Value, result.Partition, result.Offset)
		case err:=<-producer.Errors():
			log.Println("Failed to produce message", err)
		}
	}
}
