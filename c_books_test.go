package tartara

type Book struct {
	Document
	Title  string
	Author string
}

func (this *Book) SayHello() string {
	return "Hello, I am a book"
}

var Books = &Collection{
	Path:  "/tartara/db1/books/",
	Index: make(map[string]DocumentInterface),
}
