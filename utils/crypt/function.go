package crypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	PriPemBegin = "-----BEGIN RSA PRIVATE KEY-----\n"
	PriPemEnd   = "\n-----END RSA PRIVATE KEY-----"
	PubPemBegin = "-----BEGIN PUBLIC KEY-----\n"
	PubPemEnd   = "\n-----END PUBLIC KEY-----"
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

func InitPublicKey(key string, isPath bool) (pubKey *rsa.PublicKey, err error) {
	var pemData []byte

	if isPath {
		pemData, err = ioutil.ReadFile(key)
	} else {
		pemData = []byte(PubPemBegin + key + PubPemEnd)
	}
	if err != nil {
		return
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		err = errors.New("invalid public key")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err = errors.New("invalid public key")
		return
	}
	pubKey = pubInterface.(*rsa.PublicKey)
	return
}

func InitPrivateKey(key string, isPath bool) (priKey *rsa.PrivateKey, err error) {
	var pemData []byte

	if isPath {
		pemData, err = ioutil.ReadFile(key)
	} else {
		pemData = []byte(PriPemBegin + key + PriPemEnd)
	}
	fmt.Println(string(pemData))

	if err != nil {
		err = errors.New("invalid private key")
		return
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		err = errors.New("invalid private key")
		return
	}
	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		err = errors.New("invalid private key")
		return
	}

	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		err = errors.New("invalid private key error:" + err.Error())
		return
	}
	priKey = priInterface.(*rsa.PrivateKey)
	return
}
