package auth

import (
	"crypto/rsa"

	"io/ioutil"
	"log"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/orlmonteverde/go-notes-apirest/models"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("keys/private.rsa")
	if err != nil {
		log.Fatal("no se pudo leer el archivo privado", err)
	}
	publicBytes, err := ioutil.ReadFile("keys/public.rsa.pub")
	if err != nil {
		log.Fatal("no se pudo leer el archivo publico", err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse del privatekey")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse del publickey")
	}
}

func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			Issuer:    "Orlando Monteverde <orlmicron@gmail.com>",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el Token")
	}
	return result
}
