package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

const Salt = "mby"

func EncodeSha256(value string) string {
	m := sha256.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func EncodeSha256WithSalt(value string) string {

	saltedValue := Salt + value

	m := sha256.New()

	m.Write([]byte(saltedValue))

	return hex.EncodeToString(m.Sum(nil))
}

func EncodeMD5WithSalt(value string) string {
	saltedValue := Salt + value

	m := md5.New()

	m.Write([]byte(saltedValue))

	return hex.EncodeToString(m.Sum(nil))
}
