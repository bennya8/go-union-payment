package crypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func MD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func HmacSha1(text string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(text))
	return string(h.Sum(nil))
}
