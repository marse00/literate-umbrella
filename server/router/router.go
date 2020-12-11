package router

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"middleware"
	"models"
	"net/http"
)

//Handles account sign ups
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	decoder := json.NewDecoder(r.Body)
	var t models.RegisterUserData
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}
	middleware.SignUp(t.User.Username, t.User.Email, t.User.Password)
}

//Example of a GET req handling tb removed if needed
func IsLoggedIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("GET came through")
	fmt.Println(r)
	return
}

//General options/request headers
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling options...")

	if r.Header.Get("Access-Control-Request-Method") != "" {
		// Set CORS headers
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
		header.Set("Access-Control-Allow-Headers", "*")
	}

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}

//Router creation
func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", IsLoggedIn)
	router.POST("/register", Register)
	//router.POST("/login", Login)
	router.GlobalOPTIONS = http.HandlerFunc(preflightHandler)

	return router
}
