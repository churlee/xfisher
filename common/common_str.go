package common

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

var (
	r           *rand.Rand
	key               = []byte("YP_NAV_USER_SECRET")
	expiresTime int64 = 3600 * 24
)

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
//noinspection SpellCheckingInspection
func RandString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 产生json web token
func GenToken(userId string) string {
	claims := jwt.StandardClaims{
		Audience:  userId,                          // 受众
		ExpiresAt: time.Now().Unix() + expiresTime, // 失效时间
		Id:        userId,                          // 编号
		IssuedAt:  time.Now().Unix(),               // 签发时间
		Issuer:    Issuer,                          // 签发人
		NotBefore: time.Now().Unix(),               // 生效时间
		Subject:   "login",                         // 主题
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	return ss
}

func CheckToken(token string) (string, bool) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return key, nil
	})
	if err != nil || jwtToken == nil {
		log.Println("解析Token失败", err)
		return "", false
	}
	if jwtToken.Valid {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok {
			if claim.Issuer != Issuer {
				return "", false
			}
			if claim.ExpiresAt < time.Now().Unix() {
				return "", false
			}
			return claim.Id, true
		}
	}
	return "", false
}

func GenPwd(pwd, salt string) string {
	hash, _ := scrypt.Key([]byte(pwd), []byte(salt), 1<<15, 8, 1, 32)
	return fmt.Sprintf("%x", hash)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
