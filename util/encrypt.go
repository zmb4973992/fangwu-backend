package util

import (
	"golang.org/x/crypto/bcrypt"
)

type EncryptAndCompare struct{}

func (e *EncryptAndCompare) Encrypt(original string) (encrypted string, err error) {
	//第二个参数为加密难度，取值范围为4-31，官方建议10。值越大，越占用cpu
	bytes, err := bcrypt.GenerateFromPassword([]byte(original), 4)
	if err != nil {
		return "", err
	}
	encrypted = string(bytes)
	return encrypted, nil
}

func (e *EncryptAndCompare) Compare(original string, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(original))
	if err != nil {
		return false
	} else {
		return true
	}
}
