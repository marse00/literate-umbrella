module literate-umbrella

go 1.15

replace middleware => ./middleware

replace models => ./models

replace router => ./router

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	middleware v0.0.0-00010101000000-000000000000
	router v0.0.0-00010101000000-000000000000
)
