# tartara
Object oriented database


WARNING: Ultra beta prototype


## Getting started

```go
// Define a document
type Fruit struct {
	Document
	Name  string
	Color string
}

// Define a collection
var Fruits = &Collection{
	Name: "Fruits",
	Path: "/tartara/db1/fruits/",
	Index: make(map[string]DocumentInterface),
}

// Insert a document
var apple = &Fruit{
	Name:  "Apple",
	Color: "Green",
}

Fruits.Insert(apple)

// Retrieve a document
fruit := Fruits.FindById("8264535c0c4e9c6c335635c4026a8022").(*Fruit)

```
