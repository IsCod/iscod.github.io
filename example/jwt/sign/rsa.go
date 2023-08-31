package sign

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type RsaSign struct {
	publicKey  string
	privateKey string
}

func NewRsaSign(privateKey, publicKey string) RsaSign {
	return RsaSign{publicKey: publicKey, privateKey: privateKey}
}

func (s *RsaSign) Encode(data []byte) (string, error) {
	d := signData{Data: data, RegisteredClaims: jwt.RegisteredClaims{
		//Issuer:    "ADSF",
		//Subject:   "63edecb5b654030011f12ea6",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}}

	t := jwt.NewWithClaims(jwt.SigningMethodRS512, d)
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.privateKey))
	if err != nil {
		return "", err
	}

	return t.SignedString(key)
}

func (s *RsaSign) Decode(sign string) ([]byte, error) {
	out := signData{}
	_, err := jwt.ParseWithClaims(sign, &out, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(s.publicKey))
	})

	if err != nil {
		return nil, err
	}
	return out.Data, err
}
