package util

import (
	"crypto/sha1"
	"encoding/base64"
)

func GenerateShortKey(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(bs)[:6]
}
