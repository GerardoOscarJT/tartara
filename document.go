package tartara

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type DocumentInterface interface {
	SayHello()
	Set(key string, value string)
	Get(key string) string
	Inc(key string, value int)
	GenerateId() bool
	GetId() string
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

func (this *Document) GetId() string {
	return this.Id
}

func (this *Document) Inc(key string, value int) {

}

func generateUniqueId() string {

	nanoseconds := time.Now().UnixNano()
	array := []byte(strconv.FormatInt(nanoseconds, 10))
	hasher := md5.New()
	hasher.Write(array)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (this *Document) GenerateId() bool {
	if "" == this.Id {
		this.Id = generateUniqueId()
		return true
	}
	return false
}
