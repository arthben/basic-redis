// Redis locking is useful in distributed system
package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLocking struct {
	client  *redis.Client
	lockKey string
}

func (r RedisLocking) acquireLock(exp time.Duration) bool {
	// Use SETNX ( Set if Not Exists)
	lockAcquire, err := r.client.SetNX(context.Background(), r.lockKey, "1", exp).Result()
	if err != nil {
		log.Fatalf("Error while SetNX - %s", err.Error())
		return false
	}
	return lockAcquire
}

func (r RedisLocking) releaseLock() {
	r.client.Del(context.Background(), r.lockKey)
}

func main() {
	// how to run ?
	// 1. open terminal window and run ( go run locking.go) -- Success Lock
	// 2. open NEW terminal windo and run ( go run locking.go) -- Faild to lock
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	lockKey := "my_lock_key"
	expiration := time.Duration(10 * time.Second)
	redisLock := RedisLocking{
		client:  client,
		lockKey: lockKey,
	}

	if redisLock.acquireLock(expiration) {
		log.Println("Lock redis resource is succees")

		// block 10 second to simulate the resource from redis still in use for processing
		time.Sleep(expiration)
		log.Println("Work Done after 10 second")

		redisLock.releaseLock()
	} else {
		log.Println("Failed lock. Resource already locked")
	}
}
