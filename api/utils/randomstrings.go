package utils

import (
	"crypto/rand"
	"fmt"
)

func RandomFilename(max int) (s string, err error) {
	b := make([]byte, max)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}
