package tartara

import "testing"

func Test_GenerateId(t *testing.T) {
	orange := &Fruit{
		Name: "Orange",
	}

	previous_id := orange.Id
	orange.GenerateId()

	if orange.Id == previous_id {
		t.Error("GenerateId() must set attribute `Id`")
	}

}
