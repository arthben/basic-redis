package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	// check connection to redis instance with PING command
	status, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error check status redis instance %s", err.Error())
		return
	}
	fmt.Printf("Redis Instance Status: %v\n", status)

	// store string value to redis with unlimited time
	keyValue := "key1"
	neverExpired := time.Duration(0 * time.Second) // 0 = unlimited time

	err = client.Set(context.Background(), keyValue, "Basic Redis..", neverExpired).Err()
	if err != nil {
		log.Fatalf("Error set value to redis %s\n", err.Error())
		return
	}

	// read string value from redis
	strValue, err := client.Get(context.Background(), keyValue).Result()
	if err != nil {
		log.Fatalf("Error get value from redis %s\n", err.Error())
		return
	}
	fmt.Printf("String Value From Redis : %v\n", strValue)

	// store string value to redis with limited time - 2 seconds
	keyValue2 := "key2"
	exp2 := time.Duration(2 * time.Second)

	err = client.Set(context.Background(), keyValue2, "Basic Redis With Expired time", exp2).Err()
	if err != nil {
		log.Fatalf("Error set value with expired time to redis %s\n", err.Error())
		return
	}

	// read string value for key2
	strValue2, err := client.Get(context.Background(), keyValue2).Result()
	if err != nil {
		log.Fatalf("Error get value from redis %s\n", err.Error())
		return
	}
	fmt.Printf("String Value From Redis : %v\n", strValue2)

	// read string value for key2 after 2 seconds
	time.Sleep(3 * time.Second) // block until value expired after 2 second
	strValue21, err := client.Get(context.Background(), keyValue2).Result()
	if err == redis.Nil || err != nil {
		fmt.Printf("Error get value from redis after 2 seconds %s\n", err.Error())
	} else {
		fmt.Printf("String Value From Redis after 2 seconds: %v\n", strValue21)
	}

	// Store json to redis
	type TodoItem struct {
		ID   string `json:"id"`
		Todo string `json:"todo"`
	}

	// compose json
	keyID := uuid.NewString()
	todoItem := &TodoItem{
		ID:   keyID,
		Todo: "Play games",
	}
	byteJson, err := json.Marshal(todoItem)
	if err != nil {
		log.Fatalf("Error Marshal JSON to String")
		return
	}

	// Store json value
	err = client.Set(context.Background(), keyID, string(byteJson), neverExpired).Err()
	if err != nil {
		log.Fatalf("Error set JSON value to redis %s\n", err.Error())
		return
	}

	// Get json value
	strJson, err := client.Get(context.Background(), keyID).Result()
	if err == redis.Nil || err != nil {
		log.Fatalf("Error get JSON value from redis %s\n", err.Error())
		return
	}
	fmt.Printf("JSON Value from redis: %v\n", strJson)

}
