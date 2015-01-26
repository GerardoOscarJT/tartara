package tartara

import "fmt"

type DocumentInterface interface {
	SayHello()
	Set(key string, value string)
	Get(key string) string
	Inc(key string, value int)
}

type Document struct {
	Id string
}

func (this *Document) SayHello() {
	fmt.Println("Hello (I am Document)")
}

func (this *Document) Set(key string, value string) {

}

func (this *Document) Get(key string) string {
	return "---"
}

func (this *Document) Inc(key string, value int) {

}
