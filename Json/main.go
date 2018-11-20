package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	return string(data)
}

func main() {
	content := readFile("./sample.json")
	log.Print("Got file")

	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)

	log.Print("Name: ", result["name"])
}
