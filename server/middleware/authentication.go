package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("manboSecretoTemp")

func SignUp(username string, email string, password string) bool {

	cost := 3

	if !IsUsernameTaken(username) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		fmt.Println(len(string(hashedPassword)))
		if err != nil {
			fmt.Println("GenerateFromPassword error:", err)
		} else {
			err = NewUser(username, email, string(hashedPassword))
			if err != nil {
				fmt.Println("NewUser error:", err)
			} else {
				return true
			}
		}

	}
	return false
}

func SignIn(username string, password string) (string, error) {
	passwordHash := []byte(GetPasswordHash(username))
	passwordByte := []byte(password)
	result := bcrypt.CompareHashAndPassword(passwordHash, passwordByte)
	if result == nil {
		token, err := createSessionToken(username)
		if err != nil {
			fmt.Println("createSessionToken error:", err)
			return "", err
		}
		return token, nil

	}
	return "", result
}

func createSessionToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "localhost:8080",
		"sub": username,
		"exp": time.Now().Add(time.Hour)})

	return token.SignedString(jwtKey)
}
