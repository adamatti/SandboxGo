package main

import (
	"io/ioutil"
	"log"
)

func writeFile(path string, content string) {
	data := []byte(content)
	err := ioutil.WriteFile(path, data, 0777)
	if err != nil {
		log.Fatal("Error writting file: ", err)
	}
}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	return string(data)
}

func main() {
	path := "/tmp/test.txt"
	writeFile(path, "Sample file")
	content := readFile(path)
	log.Print("Content: ", content)
}
