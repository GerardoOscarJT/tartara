package tartara

import "reflect"

// type CollectionInterface interface {
// 	GetNew() CollectionInterface
// }

type Collection struct {
	Name     string
	Path     string
	Document DocumentInterface
}

func (this *Collection) Insert() DocumentInterface {

	// TODO: clone this.Document before

	// document_type := reflect.TypeOf(this.Document)

	// result := reflect.New(document_type)

	// fmt.Println(result)

	return reflect.ValueOf(this.Document).Interface().(DocumentInterface)
}
