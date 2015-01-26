package tartara

import (
	"fmt"
	"testing"
)

type User struct {
	Document
	Name    string
	Surname string
}

func define_collection() *Collection {

	return &Collection{
		Name: "Users",
		Path: "/tartara/db1/users/",
		Document: &User{
			Name:    "Fulanito",
			Surname: "Fulanitez",
		},
	}
}

func TestDefineCollectionUsers(t *testing.T) {

	Users := define_collection()

	fmt.Println("Define collection: ", Users)
}

func TestInsert(t *testing.T) {

	Users := define_collection()

	fulanito := &User{
		Name:    "Fulanito",
		Surname: "Fulanitez",
	}

	fmt.Println(fulanito)
	Users.Insert(fulanito)
	fmt.Println(fulanito)

}
