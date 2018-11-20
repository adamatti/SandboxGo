package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal("Error setting key: ", err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		log.Fatal("Error reading key: ", err)
	}
	log.Print("Key: ", val)
}
