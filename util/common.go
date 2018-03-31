package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

// create a 32-bit string by md5
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

// Generate globally unique uuid
func UUID() string {
	uuid := UUIDString()
	if "" == uuid {
		return ""
	}
	return fmt.Sprintf("%v-%v-%v-%v-%v", uuid[0:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
}

// Error return a error with format.
func Error(format string, a ...interface{}) error {
	var text = fmt.Sprintf(format, a...)
	return errors.New(text)
}
