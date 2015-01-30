package tartara

type Animal struct {
	Document
	Name    string
	Kingdom string
}

var Animals = &Collection{
	Path:  "/tartara/db1/animals/",
	Index: make(map[string]DocumentInterface),
}
