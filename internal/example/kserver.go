package example

import (
	"fmt"
	// "os"
	// "os/signal"
	"time"
	"log"
	"strconv"

	serverv3 "github.com/CBackyx/my-go-control-plane/pkg/server/v3"
	cachev3 "github.com/CBackyx/my-go-control-plane/pkg/cache/v3"

	
	// "github.com/Shopify/sarama"
)

func RunKServer(updateSignal *serverv3.UpdateSignal, cache cachev3.SnapshotCache) {
	
	// tokenMap := make(map[string]SingleRouteInfo)

	go testUpdate(updateSignal, cache)

	// config := sarama.NewConfig()
	// config.Consumer.Return.Errors = true

	// // Specify brokers address. This is default one
	// brokers := []string{"localhost:9092"}

	// // Create new consumer
	// master, err := sarama.NewConsumer(brokers, config)
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := master.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// topic := "test"
	// // How to decide partition, is it fixed value...?
	// consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	// if err != nil {
	// 	panic(err)
	// }

	// signals := make(chan os.Signal, 1)
	// signal.Notify(signals, os.Interrupt)

	// // Count how many message processed
	// msgCount := 0

	// // Get signnal for finish
	// doneCh := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case err := <-consumer.Errors():
	// 			fmt.Println(err)
	// 		case msg := <-consumer.Messages():
	// 			msgCount++
	// 			fmt.Println("Received messages", string(msg.Key), string(msg.Value))
	// 			// If the token is presented in msg.Value using Json, then parse the Json value
	// 			// if op == "Add" {

	// 			// } else if op == "Delete" {

	// 			// }

	// 		case <-signals:
	// 			fmt.Println("Interrupt is detected")
	// 			doneCh <- struct{}{}
	// 		}
	// 	}
	// }()

	// <-doneCh
	// fmt.Println("Processed", msgCount, "messages")
}

func testUpdate(updateSignal *serverv3.UpdateSignal, cache cachev3.SnapshotCache) {
	
	c := time.Tick(5 * time.Second)
    for countDown := 10; countDown > 0; countDown-- {
		<- c
		fmt.Println(countDown)
		cache.SetSnapshot("test-id", GenerateTestSnapshot(countDown))
		log.Printf("Update to " + strconv.Itoa(countDown))
		updateSignal.Clusters <- 1
		updateSignal.Routes <- 1		
	}	
}