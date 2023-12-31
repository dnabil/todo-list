package jwtauth

import (
	"github.com/golang-jwt/jwt/v4"
)

// don't forget fill in exp field in claims.
func NewToken(claims jwt.Claims, KEY []byte) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// https://github.com/dgrijalva/jwt-go/issues/65

	signed, err := token.SignedString(KEY)
	if err != nil {
		return "", err
	}
	return signed, err
}

func DecodeToken(signedToken string, ptrClaims jwt.Claims, KEY []byte) (error) {

	token, err := jwt.ParseWithClaims(signedToken, ptrClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // method used to sign the token 
		if !ok {
			// wrong signing method
			return "", ErrWrongSigningMethod
		}
		return KEY, nil
	})

	if err != nil {
		// parse failed
		return err
	}

	if !token.Valid{
		// token is not valid
		return ErrInvalidToken
	}

	return nil
}