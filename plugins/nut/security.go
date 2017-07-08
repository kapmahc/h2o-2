package nut

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"

	"github.com/astaxie/beego"
)

var (
	aesCipher cipher.Block
	hmacKey   []byte
)

// Encrypt encrypt
func Encrypt(buf []byte) ([]byte, error) {

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(aesCipher, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

// Decrypt decrypt
func Decrypt(buf []byte) ([]byte, error) {

	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(aesCipher, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

// Sum sum hmac
func Sum(plain []byte) []byte {
	mac := hmac.New(sha512.New, hmacKey)
	return mac.Sum(plain)
}

// Chk chk hmac
func Chk(plain, code []byte) bool {
	return hmac.Equal(Sum(plain), code)
}

func init() {
	var err error
	aesCipher, err = aes.NewCipher([]byte(beego.AppConfig.String("aeskey")))
	if err != nil {
		beego.Error(err)
	}
	hmacKey = []byte(beego.AppConfig.String("hmackey"))
}
