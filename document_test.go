package tartara

import "testing"

type User struct {
	Document
	Name    string
	Surname string
}

var Users = &Collection{
	Name:  "Users",
	Path:  "/tartara/db1/users/",
	Index: make(map[string]DocumentInterface),
}

type Book struct {
	Document
	Title  string
	Author string
}

func (this *Book) SayHello() string {
	return "Hello, I am a book"
}

var Books = &Collection{
	Name:  "Books",
	Path:  "/tartara/db1/books/",
	Index: make(map[string]DocumentInterface),
}

func TestSayHello(t *testing.T) {

	fulanito := &User{
		Name:    "Fulanito",
		Surname: "Fulanitez",
	}

	if "Hello (I am Document)" != fulanito.SayHello() {
		t.Error("A `User` must return the default SayHello() implementation")
	}

}

func TestSayHelloOverride(t *testing.T) {

	jekyll_and_mister_hyde := &Book{
		Title:  "Jekyll and Mister Hyde",
		Author: "Robert Louis Stevenson",
	}

	if "Hello, I am a book" != jekyll_and_mister_hyde.SayHello() {
		t.Error("A `Book` must return a custom SayHello() implementation")
	}

}
