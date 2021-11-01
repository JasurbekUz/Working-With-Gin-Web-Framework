package utils

import (
	"app/models"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func CodeGenerator() string {
	code := make([]byte, 4)
	rand.Read(code)
	return fmt.Sprintf("%x", code)
}

var jwtKey = []byte("FizzBuzz")

func JwtTokenMaker(userName string) string {

	expirationTime := time.Now().Add(5 * time.Second)

	payload := models.Payload{

		UserName: userName,
		Role:     1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%T", tokenString)

	return tokenString

	// tekshirish

	/*tkn, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			log.Println("invalid token 1")

			return
		}

		log.Println("invalid token 2")

		return
	}

	if !tkn.Valid {

		log.Println("invalid token 3")
	}

	log.Printf("\n\n\n%T \n\n\n%v", tkn, tkn)*/
}
