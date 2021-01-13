package middleware

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(username string, email string, password string) bool {

	cost := 3

	if !IsUsernameTaken(username) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		fmt.Println(len(string(hashedPassword)))
		if err != nil {
			fmt.Println("GenerateFromPassword error:", err)
		} else {
			err = NewUser(username, email, hashedPassword)
			if err != nil {
				fmt.Println("NewUser error:", err)
			} else {
				return true
			}
		}

	}
	return false
}
