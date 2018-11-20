package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

func main() {
	db := pg.Connect(&pg.Options{
		User: "sample",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		log.Panic(err)
	}

	user1 := addUser(db)
	user1 = findUser(db, user1.Id)
	allUsers(db)
}

func addUser(db *pg.DB) *User {
	log.Print("Add user")
	user1 := &User{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	err := db.Insert(user1)
	if err != nil {
		log.Panic(err)
	}
	return user1
}

func findUser(db *pg.DB, id int64) *User {
	log.Print("Find user")

	user := &User{Id: id}
	err := db.Select(user)
	if err != nil {
		log.Panic(err)
	}
	return user
}

func allUsers(db *pg.DB) []User {
	log.Print("All users")

	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		log.Panic(err)
	}
	return users
}

func createSchema(db *pg.DB) error {
	models := []interface{}{(*User)(nil)}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{Temp: true})
		if err != nil {
			return err
		}
	}
	return nil
}
