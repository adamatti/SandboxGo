package main

import (
	"log"
	"time"

	"github.com/chebyrash/promise"
)

func main() {
	var promise1 = promise.New(func(resolve func(interface{}), reject func(error)) {
		time.Sleep(3 * time.Second)
		log.Println("done with one")
		resolve("done with one")
	})

	var promise2 = promise.New(func(resolve func(interface{}), reject func(error)) {
		time.Sleep(5 * time.Second)
		log.Println("done with two")
		resolve("done with two")
	}).Then(func(data interface{}) interface{} {
		log.Println("done with third")
		return nil
	})

	promise.AwaitAll(promise1, promise2)
	log.Println("All done")
}
