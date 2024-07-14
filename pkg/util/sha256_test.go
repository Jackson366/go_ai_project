package util

import "testing"

func TestEncodeSha256WithSalt(t *testing.T) {
	salt := EncodeSha256WithSalt("12345678")
	t.Log(salt)
}
