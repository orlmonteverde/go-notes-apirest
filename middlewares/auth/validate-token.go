package auth

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/orlmonteverde/go-notes-apirest/models"
)

func validateToken(w http.ResponseWriter, r *http.Request) (role string, err error) {
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				err = errors.New("Su token ha expirado")
				return
			case jwt.ValidationErrorSignatureInvalid:
				err = errors.New("La firma del Token no coincide")
				return
			}
		default:
			err = errors.New("Su token no es válido")
			return
		}
	}

	if claims, ok := token.Claims.(*models.Claim); token.Valid && ok {
		role = claims.Role
		return
	} else {
		err = errors.New("No autorizado")
		return
	}

}
