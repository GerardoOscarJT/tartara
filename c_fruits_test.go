package tartara

type Fruit struct {
	Document
	Name  string
	Color string
}

var Fruits = &Collection{
	Path:  "/tartara/db1/fruits/",
	Index: make(map[string]DocumentInterface),
}
