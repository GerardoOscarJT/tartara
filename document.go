package tartara

import "github.com/GerardoOscarJT/tartara/utils"

type DocumentInterface interface {
	SayHello() string
	Set(key string, value string)
	Get(key string) string
	Inc(key string, value int)
	GenerateId() bool
	GetId() string
}

type Document struct {
	Id string
}

func (this *Document) SayHello() string {
	return "Hello (I am Document)"
}

func (this *Document) Set(key string, value string) {

}

func (this *Document) Get(key string) string {
	return "---"
}

func (this *Document) GetId() string {
	return this.Id
}

func (this *Document) Inc(key string, value int) {

}

func (this *Document) GenerateId() bool {
	if "" == this.Id {
		this.Id = utils.GenerateUniqueId()
		return true
	}
	return false
}
