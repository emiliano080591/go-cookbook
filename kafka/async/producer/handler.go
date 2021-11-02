package producer

import (
	"gopkg.in/Shopify/sarama.v1"
	"net/http"
)

type KafkaController struct {
	Producer sarama.AsyncProducer
}

func (k *KafkaController) Handler(w http.ResponseWriter,r *http.Request) {
	if err:=r.ParseForm();err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	msg:=r.FormValue("msg")
	if msg=="" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("msg must be set"))
		return
	}
	k.Producer.Input() <- &sarama.ProducerMessage{
		Topic: "example",
		Key: nil,
		Value: sarama.StringEncoder(r.FormValue("msg")),
	}
	w.WriteHeader(http.StatusOK)
}
