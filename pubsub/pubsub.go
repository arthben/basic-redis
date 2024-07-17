package main

import (
	"context"
	"flag"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// How to run:
	// 1. Open terminal window ( as subscriber ), type : go run pubsub.go --mode=subscriber
	// 2. Open NEW terminal window ( as publisher ), type : go run pubsub.go --mode=publisher
	mode := flag.String("mode", "publisher", "Mode: publisher / subscriber. Default mode is publisher")
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	channel := "basic-redis-pubsub"
	switch *mode {
	case "publisher":
		publishMessage(client, channel)

	case "subscriber":
		subscribe(client, channel)
	default:
		log.Println("Invalid mode. Supported mode : publisher / subscriber")
		return
	}
}

func subscribe(client *redis.Client, channel string) {
	subs := client.Subscribe(context.Background(), channel)
	for {
		msg, err := subs.ReceiveMessage(context.Background())
		if err != nil {
			log.Fatalf("Subscriber failed receive message - %s", err.Error())
			break
		}
		log.Printf("Receive Message : %s - From Channel %s", msg.Payload, msg.Channel)
	}
}

func publishMessage(client *redis.Client, channel string) {
	message := "Hai.. I'm publisher"
	err := client.Publish(context.Background(), channel, message).Err()
	if err != nil {
		log.Fatalf("Failed publish. Error - %s", err.Error())
	} else {
		log.Printf("Published Message : %s", message)
	}
}
