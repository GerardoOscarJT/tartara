package tartara

import (
	"fmt"
	"testing"
	"time"
)

// Define collection `Fruits` //////////////////////

type Fruit struct {
	Document
	Name  string
	Color string
}

var Fruits = &Collection{
	Name:  "Fruits",
	Path:  "/tartara/db1/fruits/",
	Index: make(map[string]DocumentInterface),
}

////////////////////////////////////////////////////

// Define collection `Animals` //////////////////////

type Animal struct {
	Document
	Name    string
	Kingdom string
}

var Animals = &Collection{
	Name:  "Animals",
	Path:  "/tartara/db1/animals/",
	Index: make(map[string]DocumentInterface),
}

////////////////////////////////////////////////////

func TestInsert(t *testing.T) {
	var apple = &Fruit{
		Name:  "Apple",
		Color: "Green",
	}

	Fruits.Insert(apple)

	// TODO: Check if the document has been stored
}

func BenchmarkInsert(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {

		var fruit = &Fruit{
			Name:  "Fruit" + string(i),
			Color: "#" + string(i),
		}

		_ = fruit

		Fruits.Insert(fruit)
	}
	end := time.Now()

	fmt.Println("N:", b.N, "\t", "Time (ms):", (end.Nanosecond()-start.Nanosecond())/1000000)
}

func TestGetById(t *testing.T) {
	var apple = &Fruit{
		Name:  "Apple",
		Color: "Red",
	}

	Fruits.Insert(apple)

	fruit := Fruits.FindById(apple.GetId()).(*Fruit)

	if apple.GetId() != fruit.GetId() {
		t.Error("`Id` does not match")
	}

	if apple.Name != fruit.Name {
		t.Error("`Name` does not matach")
	}
}

func define_collection_fruits() *Collection {
	return &Collection{
		Name:  "Fruits",
		Path:  "/tartara/db1/fruits/",
		Index: make(map[string]DocumentInterface),
	}
}
