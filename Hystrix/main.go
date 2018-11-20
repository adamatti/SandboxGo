package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gorilla/mux"
)

func enableMetrics() {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "3001"), hystrixStreamHandler)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func startHttp() {
	router := mux.NewRouter()

	router.HandleFunc("/api", func(rw http.ResponseWriter, r *http.Request) {
		hystrix.Do("my_command", func() error {
			// talk to other services
			if randInt(0, 3) != 1 {
				fmt.Fprintf(rw, "Working")
				return nil
			}
			return errors.New("generated error")
		}, func(err error) error {
			// do this when services are down
			fmt.Fprintf(rw, "Not working")
			return nil
		})
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	go enableMetrics()
	startHttp()
}
