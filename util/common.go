package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// create a 32-bit string by md5.
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// create a uuid string
func UUIDString() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(getMd5String(base64.URLEncoding.EncodeToString(b)))
}

// Generate globally unique uuid.
func UUID() string {
	uuid := UUIDString()
	if "" == uuid {
		return ""
	}
	return fmt.Sprintf("%v-%v-%v-%v-%v", uuid[0:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
}

// NewError return a error with format.
func NewError(format string, a ...interface{}) error {
	var text = fmt.Sprintf(format, a...)
	return errors.New(text)
}

// Now13 return 13 bit timestamp.
func Now13() int64 {
	var now = time.Now()
	return int64(now.UnixNano() / 1e6)
}

// Now return 10 bit timestamp.
func Now10() int64 {
	var now = time.Now()
	return now.Unix()
}

// S2Json trans data to json, e.g struct, map and so on.
func S2Json(data interface{}) string {
	bys, _ := json.Marshal(data)
	return string(bys)
}

// Json2S trans json to object
func Json2S(src string, dest interface{}) error {
	return json.Unmarshal([]byte(src), dest)
}
