package tartara

import "github.com/GerardoOscarJT/tartara/utils"

type DocumentInterface interface {
	Set(key string, value string)
	Get(key string) string
	Inc(key string, value int)
	GenerateId() bool
	GetId() string
}

type Document struct {
	Id string
}

func (this *Document) GetId() string {
	return this.Id
}

func (this *Document) GenerateId() bool {
	if "" == this.Id {
		this.Id = utils.GenerateUniqueId()
		return true
	}
	return false
}

func (this *Document) Set(key string, value string) {
	// TODO: implement this
}

func (this *Document) Get(key string) string {
	// TODO: implement this
	return "---"
}

func (this *Document) Inc(key string, value int) {
	// TODO: implement this
}
