package main

import (
	"github.com/agtorre/go-cookbook/chapter11/kafkaflow"
	flow "github.com/trustmaster/goflow"
	sarama "gopkg.in/Shopify/sarama.v1"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	net := kafkaflow.NewUpperApp()

	in := make(chan string)
	net.SetInPort("In", in)

	flow.RunNet(net)
	defer func() {
		close(in)
		<-net.Wait()
	}()

	for {
		msg := <-partitionConsumer.Messages()
		in <- string(msg.Value)
	}

}
