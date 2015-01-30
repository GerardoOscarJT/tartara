package tartara

import "testing"

func TestInsert(t *testing.T) {
	var apple = &Fruit{
		Name:  "Apple",
		Color: "Green",
	}

	Fruits.Insert(apple)

	// TODO: Check if the document has been stored
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
