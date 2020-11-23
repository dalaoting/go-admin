package uuid

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Test_uuid(t *testing.T) {
	//
	// ret, err := base64.StdEncoding.DecodeString("eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLmltYmFjay5wYWxwYWwiLCJleHAiOjE1ODcwMTgzNjYsImlhdCI6MTU4NzAxNzc2Niwic3ViIjoiMDAxMjc2LjczNGUyYWI0NTZkNTRiNDE4MjNkN2QyMTIzNzllMzE0LjE0NDEiLCJjX2hhc2giOiJFZkZaRDdJMGFUSU1oNUt1cHR0ZWhBIiwiZW1haWwiOiJqc2g2OXM1OGpiQHByaXZhdGVyZWxheS5hcHBsZWlkLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjoidHJ1ZSIsImlzX3ByaXZhdGVfZW1haWwiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNTg3MDE3NzY2LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.HF4m-sj44VnfZrsMfi2981B1C9CbEUygzR_aAHptBmkXTDCCZrh44VPmx8IIby56uWn52Xt4d_Cb-Xx8fzyxFAX_BZgxDexEM9oV3ILT0NCQvIe91cYzUV6YlFwOsIwjHDztq7rGQAPXgh04QRRIQFnoWRgSLzEliSBKWDhI_mozSctgbgZYG7AFSuSabRNBpirkIcO8HATx8KO6fHbGTYMp_bnWc70lPXu7bM0046OpvRkNNi0b2nabRUfI1Vi7AcykdENr0zKuFxcssx4hfnTszGqJDDaoTW6jZ23B-CTCN7ONHy1C7ubWteUJUmAWvmkbnSAiRfeLWGJ40fClzQ")
	// if err != nil {
	// 	println(err.Error())
	// }
	// println(string(ret))

	// var tokenString = "eyJraWQiOiJlWGF1bm1MIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLmltYmFjay5wYWxwYWwiLCJleHAiOjE1ODcwMjE4NjQsImlhdCI6MTU4NzAyMTI2NCwic3ViIjoiMDAwODAxLjFjYTE1MGMwZTI0YzRmODU5Y2YzOTE4NjEzZDZiOTI4LjAzMzUiLCJjX2hhc2giOiJ6STVweGotOTNxa0ZHNjAyQk96TnRBIiwiYXV0aF90aW1lIjoxNTg3MDIxMjY0LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.FpvIkuy9PQIgqimYjNzFfPoX6tLaF2IwsexuNu2MK5mmtWAaBV0JisTQlQXV5OySXmrxglmTuwEljk7xmQ4DAa8tvGS1ZY_U112zW_ZH2HlSH8NJA1OH2HwfxWzWx0cxsIRupmGIuslfxnHNVce7EpqdsHpyyi_wQuc5oVBytkGWaXCIgwJ9zc5veS_1IITeXxtekECrH9ECI1mRmM4BW0POfZ12IBLtMVjC6ZgyaISmsqnQLc4pJRX5YKXZS0Ivt1zgkTT_K4n7USdJt6uJCSBA3KtPM76KJp2fuk_lpxzhn95pmM-3QyUsO7YnN0LYOvPMFlZ8yvrWrxwmEn_ccw"

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "W6HUP5T429",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 180).Unix(),
		"aud": "https://appleid.apple.com",
		"sub": "com.imback.palpal",
	})
	token.Header["alg"] = "ES256"
	token.Header["kid"] = "8KHJG3QM4M"

	// token.Method = jwt.SigningMethodES256

	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgkYH4mf56TCG1lcrI
svdMlMN8nLmfRaOowpzNi5Z+mV2gCgYIKoZIzj0DAQehRANCAATE/oWwlKP7RNho
aCpnA2oi6nAZXadbuao488Yii+C/vdHSJfDdWB3ebwWaZwgUKxkJlARVz1XlJhqS
PIAqec5q
-----END PRIVATE KEY-----`))
	if err != nil {
		println(err.Error())
		return
	}
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)
	if err != nil {
		println(err.Error())
		return
	}

	println(tokenString)

	InitUUID(time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local))
	t.Log(UUID())
}

func BenchmarkUUID(b *testing.B) {
	InitUUID()

	b.StartTimer()
	// for i := 0; i < 10000; i++ {
	UUID()
	// }
	b.StopTimer()
}

func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("Invalid Key: Key must be PEM encoded PKCS1 or PKCS8 private key")
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	switch v := parsedKey.(type) {
	default:
		println(v)
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, errors.New("Key is not a valid RSA private key")
	}

	return pkey, nil
}
