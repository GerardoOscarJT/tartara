package tartara

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/GerardoOscarJT/tartara/utils"
)

const (
	MOD = os.FileMode(0744)
)

type CollectionInterface interface {
	FindById(id string) DocumentInterface
	Insert(document DocumentInterface) bool
}

type Collection struct {
	Name     string
	Path     string
	Document DocumentInterface
	Index    map[string]DocumentInterface
}

func (this *Collection) FindById(id string) DocumentInterface {
	return this.Index[id]
}

func (this *Collection) Insert(document DocumentInterface) bool {

	if !document.GenerateId() {
		fmt.Println("Error generating document id")
		return false
	}

	if nil == this.store_document(document) {
		return true
	}

	if nil != os.MkdirAll(this.Path, MOD) {
		fmt.Println("Can not create path", this.Path)
		return false
	}

	// Retry document insertion
	if nil == this.store_document(document) {
		return true
	}

	fmt.Println("Document could not be inserted.")
	return false
}

func (this *Collection) update_index(document DocumentInterface) {
	this.Index[document.GetId()] = document
}

func (this *Collection) store_document(document DocumentInterface) error {
	path := this.Path + document.GetId()
	text := utils.Document2json(document)

	err := ioutil.WriteFile(path, text, MOD)

	if nil == err {
		this.update_index(document)
	}

	return err
}
