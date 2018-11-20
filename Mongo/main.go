package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("sample")
	c := db.C("person")

	insert(c)
	find(c)
}

func insert(c *mgo.Collection) {
	err := c.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}
}

func find(c *mgo.Collection) {
	result := Person{}
	err := c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Phone:", result.Phone)
}
