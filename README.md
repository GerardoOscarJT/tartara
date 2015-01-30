# tartara
Object oriented database


WARNING: Ultra beta prototype


## Getting started

```go
// Define a document:
type Fruit struct {
	Document
	Name  string
	Color string
}

// Define a collection:
var Fruits = &Collection{
	Name: "Fruits",
	Path: "/tartara/db1/fruits/",
	Index: make(map[string]DocumentInterface),
}

// Insert a document:
var apple = &Fruit{
	Name:  "Apple",
	Color: "Green",
}

Fruits.Insert(apple)

// Retrieve a document:
fruit := Fruits.FindById("8264535c0c4e9c6c335635c4026a8022").(*Fruit)

```

## Extending collection

Good to avoid type cast: <code>Users.Find("j9983")</code> instead of
<code>Users.Find("j9983").(*User)</code>

```go
type User struct {
	Document
	Nick        string
	Email       string
	Password    string
}

type UsersCollection struct {
	Collection
}

// Extension:
func (this *UsersCollection) FindById(id string) *User {
	return Users.Collection.FindById(id).(*User)
}

// Instanciate collection:
var Users = UsersCollection{
	Collection{
		Path:  "/tartara/db1/users/",
		Index: make(map[string]DocumentInterface),
	},
}


```