package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

func GenerateUniqueId() string {

	nanoseconds := time.Now().UnixNano()
	array := []byte(strconv.FormatInt(nanoseconds, 10))
	hasher := md5.New()
	hasher.Write(array)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Document2json(document interface{}) []byte {
	j, err := json.Marshal(document)
	if err != nil {
		return []byte{}
	}
	return j
}
