module router

go 1.15

replace middleware => ../middleware

replace models => ../models

require (
	github.com/julienschmidt/httprouter v1.3.0
	middleware v0.0.0-00010101000000-000000000000
	models v0.0.0-00010101000000-000000000000
)
