package main

import (
	"log"
	"math"
	"strconv"
)

func main() {
	log.Print("Started")
	for i := 1; i <= 15; i++ {
		log.Print(i, " -> ", fizzBuzz(float64(i)))
	}
}

func fizzBuzz(x float64) string {
	if math.Mod(x, 3) == 0.0 && math.Mod(x, 5) == 0.0 {
		return "FizzBuzz"
	}
	if math.Mod(x, 3) == 0.0 {
		return "Fizz"
	}
	if math.Mod(x, 5) == 0.0 {
		return "Buzz"
	}
	return strconv.FormatFloat(x, 'f', 6, 64)
}
