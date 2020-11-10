/*
INFO:
	Currently not working with Firefox, researching on that
*/

package router

import (
	"log"
	"fmt"
	"github.com/julienschmidt/httprouter"
	//"middleware"
	"encoding/json"
	"net/http"
)

type loginDetails struct {
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
 }

//Example of a POST req handling tb removed if needed
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("login POST came through")

	decoder := json.NewDecoder(r.Body)
	var t loginDetails
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(r.Body, " --> POST before json decoding")
	fmt.Println(t, " --> POST after json decoding")

	fmt.Println(t.User.Username)
	fmt.Println(t.User.Password)

	//middleware.SetPassword(t.User.Username, t.User.Password)

	return
}

//Example of a GET req handling tb removed if needed
func IsLoggedIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("GET came through");
	fmt.Println(r);
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
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*" )
	}

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}

//Router creation
func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/",IsLoggedIn)
	router.POST("/",Login)
	router.GlobalOPTIONS = http.HandlerFunc(preflightHandler)

	return router
}
