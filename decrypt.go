package main

import (
	"fmt"
	s "strings"
	p "golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
	"net/url"
	"github.com/spacemonkeygo/openssl"
	"encoding/base64"
)

const (
	salt = "authenticated encrypted cookie"
)

type enc_cookie struct {
	body []byte
	iv []byte
	auth_tag []byte
}

func get_enc_cookie(raw string) (ec enc_cookie, err error){
	var q enc_cookie
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
	return enc_cookie{body: body, iv: iv, auth_tag: auth_tag }, nil
}

func Rails(cookie string, secret_key_base []byte) {
	message, err := get_enc_cookie(cookie)
	secret := p.Key(secret_key_base, []byte(salt), 1000, 32, sha1.New) 
	ctx, err := openssl.NewGCMDecryptionCipherCtx(len(secret)*8, nil, secret, message.iv)
	if err!= nil {
		fmt.Println("hello")
	}

	data, err := ctx.DecryptUpdate(message.body)
	fmt.Println(err)
	fmt.Println(string(data))}

func main() {
	cookie := "6g0bK4or%2FsOhFzOiOEkwQ9YneH1wvHMfdqH5ZsFlhU8qDNTqC4hZ7h%2FjxUu6Q%2FHxSjRnDIl3XFql8JQC25XWWJDYiI%2F7enasOZAAoC7xVVqRVbuJip7LKBnp9he511kOMcSruaID2Hl%2BO8QJEyz2B99OHL3wRoUKep7upjO5dPfOhmCJtiAAZdoKRHnE0SA6FG9bI%2FWLVMnvFAL7yLKaHYdR1spcs6NJ%2FTIWZWs0tYMSGC1c6PoOzsXY94fiykZEJ8%2F9yizWMlVwb1F%2Fl94E8cKxNjTeasUqCRx2hQx6Jit6srPyX6qm6cQVx0bnk2zieA%3D%3D--lYGRqA1Q707vn8rp--RKgU59EtqNE%2BKQlu4cgmuA%3D%3D"
	secret_key_base := []byte("ef0eee90f56ce0149fd9bce4a014aa239bfdb3dfac6c9428a86b5b1a93f217675e471fe8066265bd662b7f2dad76ba77ed1bb69dcebde439915c8e8ae4cb6e69")
	Rails(cookie, secret_key_base)
	fmt.Println("testing")
}