package util

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1加密
func Sha1(str string) string {

	h := sha1.New()
	
	h.Write([]byte(str))

	result := h.Sum(nil)

	return hex.EncodeToString(result)
}
