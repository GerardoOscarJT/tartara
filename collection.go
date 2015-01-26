package tartara

import (
	"encoding/json"
	"io/ioutil"
)

// type CollectionInterface interface {
// 	GetNew() CollectionInterface
// }

type Collection struct {
	Name     string
	Path     string
	Document DocumentInterface
}

func (this *Collection) Insert(document DocumentInterface) bool {

	if document.GenerateId() {
		err := ioutil.WriteFile(this.Path+"/"+document.GetId(), doc2json(document), 0644)
		_ = err
		return true
	}

	return false
}

func doc2json(document interface{}) []byte {
	j, err := json.Marshal(document)
	if err != nil {
		return []byte{}
	}
	return j
}
