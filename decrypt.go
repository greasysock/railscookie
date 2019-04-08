package railscookie

import (
	"fmt"
	s "strings"
	p "golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
	"net/url"
	"net/http"
	"github.com/spacemonkeygo/openssl"
	"encoding/base64"
)

type EncCookie struct {
	body []byte
	iv []byte
	auth_tag []byte
}

func getEncCookie(raw string) (ec EncCookie, err error){
	var q EncCookie
	// First decode url safe message to url unsafe/original message
	url_unsafe, err := url.QueryUnescape(raw)

	if err != nil{
		return q, err
	}

	//Split stuff up into message/init vector/auth
	stuff := s.Split(url_unsafe, "--")
	//Decode the base64 body to original message
	body, err := base64.StdEncoding.DecodeString(stuff[0])
	if err != nil{return q, err}
	iv, err := base64.StdEncoding.DecodeString(stuff[1])
	if err != nil{return q, err}
	auth_tag_before := stuff[2]
	auth_tag := []byte(auth_tag_before)
	return EncCookie{body: body, iv: iv, auth_tag: auth_tag }, nil
}

func verify(cookie EncCookie) {}

func decrypt(cookie EncCookie) {}

func checkCookieValid(cookie *http.Cookie) (err error){
	return nil
}

func DecryptAndVerify(cookie *http.Cookie, secret_key_base []byte) {
	err := checkCookieValid(cookie)
	if err != nil {return}
	message, err := getEncCookie(cookie.Value)
	if err != nil {return}
	secret := p.Key(secret_key_base, []byte(salt), 1000, 32, sha1.New) 
	ctx, err := openssl.NewGCMDecryptionCipherCtx(len(secret)*8, nil, secret, message.iv)
	if err!= nil {return}

	data, err := ctx.DecryptUpdate(message.body)
	fmt.Println(string(data))}