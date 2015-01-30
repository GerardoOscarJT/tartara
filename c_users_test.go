package tartara

/***
The main target of this set of tests is to explore different ways of extending
the document and the collection.


Extend collections is useful to provide semantic to the user of the library. For
example, to rewrite `FindById` to avoid type casting:

	Users.FindById("8352774d")

instead of:

	Users.FindById("8352774d").(Users)


Extending documents to add new behaviour but preserving the previous one. For
example, add a password generation feature:

	Users.FindById("8352774d").GeneratePassword()


Conclusions (to now):
-   Extending the document is easy: attach a function to it
-   Extending the collection is triky: create a struct inherited by `Collection`
    and then attach a function to it

***/

import (
	"reflect"
	"testing"

	"github.com/GerardoOscarJT/tartara/utils"
)

// Definitions:

type User struct {
	Document
	Nick        string
	Email       string
	Password    string
	YearOfBirth int
}

type UsersCollection struct {
	Collection
}

// Extension:
func (this *UsersCollection) FindById(id string) *User {
	return Users.Collection.FindById(id).(*User)
}

func (this *User) GeneratePassword() {
	this.Password = utils.GenerateUniqueId()
}

// Instanciation:
var Users = UsersCollection{
	Collection{
		Path:  "/tartara/db1/users/",
		Index: make(map[string]DocumentInterface),
	},
}

// Tests:
func TestOverrideFindById(t *testing.T) {
	john := &User{
		Nick:        "John",
		Email:       "john@gmail.com",
		Password:    "",
		YearOfBirth: 2000,
	}

	Users.Insert(john)

	user := Users.FindById(john.Id)

	if reflect.TypeOf(user) != reflect.TypeOf(john) {
		t.Error("Type must be equal")
	}

}

func TestExtendsFunctionality(t *testing.T) {
	john := &User{
		Nick: "John",
	}

	john.GeneratePassword()
	first := john.Password

	john.GeneratePassword()
	second := john.Password

	if first == second {
		t.Error("GeneratePassword() must return different passwords")
	}
}
