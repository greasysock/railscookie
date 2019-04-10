# railscookie

This library is designed to simply encrypt and decrypt Rails 5 cookies, specifically designed for Rails 5.2.2

To import:

`import 	"github.com/greasysock/railscookie"`

If you plan to store your secret_key_base, please make sure that it is in an external location and not your compiled application because they will not be safe. I suggest putting it into an enviornment variable.

## Interface

`DecryptAndVerify(cookie *http.Cookie, secret_key_base []byte) (raw_cookie []byte, err error)`

Decrypt and Verify does exactly what it says, it will decrypt the cookie and verify it. It will either return an error or a raw decrypted byte array of the cookie for further decoding with json, etc.

### TODO

* Encrypt
* Verify signature
* Create signature for encrypted Cookie.
