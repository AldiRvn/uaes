package uaes

import (
	"encoding/json"
	"log"

	"github.com/Luzifer/go-openssl/v4"
)

type Aes struct {
	secretKey string
	cg        openssl.CredsGenerator
}

// NewAES Constructor for Aes
func NewAES(secretKey string) *Aes {
	o := new(Aes)
	o.secretKey = secretKey

	o.cg = openssl.BytesToKeyMD5 //? Compatible with Crypto JS
	return o
}

func (uaes *Aes) Decrypt(enc string) (res string) {
	dec, err := openssl.New().DecryptBytes(uaes.secretKey, []byte(enc), uaes.cg)
	if err != nil {
		log.Println(err, enc)
		return enc //? Return original
	}

	res = string(dec)
	return
}
func (uaes *Aes) DecryptToMap(enc string) (res map[string]any) {
	if err := json.Unmarshal([]byte(uaes.Decrypt(enc)), &res); err != nil {
		log.Println(err)
		return
	}
	return
}
func (uaes *Aes) Encrypt(target []byte) (res string, err error) {
	resRaw, err := openssl.New().EncryptBytes(uaes.secretKey, target, uaes.cg)
	if err != nil {
		log.Println(err)
		return
	}
	res = string(resRaw)
	return
}
func (uaes *Aes) EncryptAny(target any) (res string, err error) {
	asJson, err := json.Marshal(target)
	if err != nil {
		log.Println(err)
		return
	}

	resRaw, err := openssl.New().EncryptBytes(uaes.secretKey, asJson, uaes.cg)
	if err != nil {
		log.Println(err)
		return
	}
	res = string(resRaw)
	return
}
