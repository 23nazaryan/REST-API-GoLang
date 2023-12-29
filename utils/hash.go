package utils

import (
	"crypto/rand"
	"encoding/base32"
	"strings"
)

func RandomToken() string {
	rnd := make([]byte, 64)

	_, err := rand.Read(rnd)
	if err != nil {
		panic(err)
	}

	return strings.Trim(base32.StdEncoding.EncodeToString(rnd), "=")
}
