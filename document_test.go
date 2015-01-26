package tartara

import (
	"fmt"
	"testing"
)

type User struct {
	Document
	Name string
}

func define_collection() *Collection {

	return &Collection{
		Name: "Users",
		Path: "/tartara/db1/users/",
		Document: &User{
			Name: "Fulanito",
		},
	}
}

func TestDefineCollectionUsers(t *testing.T) {

	Users := define_collection()

	fmt.Println("Define collection: ", Users)
}

func TestInsert(t *testing.T) {

	Users := define_collection()
	fmt.Println(Users)

	// a := Users.Insert()
	// fmt.Println(a)

	// b := Users.Insert()
	// fmt.Println(b)

	// b.(User).Name = "MenGANO"

	// fmt.Println(a)
	// fmt.Println(b)

}
