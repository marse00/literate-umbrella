package router

import (
	"encoding/json"
	"fmt"
	"log"
	"middleware"
	"models"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
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

//Handles account sign ins
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t models.LoginUserData
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}
	JWToken, err := middleware.SignIn(strings.TrimSpace(t.User.Username), strings.TrimSpace(t.User.Password))
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	header.Set("username", strings.TrimSpace(t.User.Username))
	http.SetCookie(w, &http.Cookie{Name: "userToken", Value: string(JWToken), Path: "/", HttpOnly: true, SameSite: http.SameSiteNoneMode, Secure: false, Domain: "localhost"})
	w.WriteHeader(http.StatusOK)
}

//Example of a GET req handling tb removed if needed
func IsLoggedIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("GET came through")

	return
}

func VideoPlayerGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Will get hash key from url (?hash=<hash-key>) and send it to the decoder that will then save the id of the video
	return
}

func VideoPlayerPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Will use the decripted id from video and get the video from the database to then use it in the post request
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
		header.Set("Access-Control-Allow-Headers", "Content-Type, withcredentials")
		header.Set("Access-Control-Allow-Credentials", "true")
	}

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}

//Router creation
func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/", Login)
	router.GET("/watch", VideoPlayerGet)
	router.POST("/watch", VideoPlayerPost)
	router.GET("/", IsLoggedIn)
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.OPTIONS("/login", Login)
	//router.POST("/login", Login)
	//router.GlobalOPTIONS = http.HandlerFunc(preflightHandler)

	return router
}
