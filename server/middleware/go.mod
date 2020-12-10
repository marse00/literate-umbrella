module middleware

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.1.2
	github.com/lib/pq v1.8.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
)

replace models => ../models