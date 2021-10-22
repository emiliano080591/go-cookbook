package main

import (
	"context"
	"fmt"
	"github.com/emiliano080591/go-cookbook/pipeline"
	"time"
)

func main()  {
	ctx:=context.Background()
	ctx,cancel:=context.WithTimeout(ctx,2*time.Second)
	defer cancel()

	in,out:=pipeline.NewPipeline(ctx,10,10)

	go func() {
		for i := 0; i <10 ; i++ {
			in<-fmt.Sprint("Message",i)
		}
	}()

	for i := 0; i < 10; i++ {
		<-out
	}
}
