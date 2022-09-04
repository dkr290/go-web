package main

import (
	"log"
	"time"
)

// might need this user varible to be available in anothr package to be exported in another package
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		FirstName: "Trevor",
		LastName:  "Bond",
		Age:       21,
	}

	log.Println(user.FirstName, user.LastName, user.Age, "BirthDate:", user.BirthDate)

}
