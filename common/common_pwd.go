package common

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func ValidatePwd(pwd, rpwd, salt string) bool {
	hash, err := scrypt.Key([]byte(rpwd), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return false
	}
	str := fmt.Sprintf("%x", hash)
	return pwd == str
}
