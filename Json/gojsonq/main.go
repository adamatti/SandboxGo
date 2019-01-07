package main

import "github.com/thedevsaddam/gojsonq"

const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`

func main() {
	name := gojsonq.New().JSONString(json).Find("name.first")
	println(name.(string)) // Tom
}
