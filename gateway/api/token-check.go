package api

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

// this example in jwtauth...I dont really understand how it fits in with jwt-go library
var tokenAuth *jwtauth.JWTAuth

// this handles verifying token 
func tokenCheckrouter() http.Handler {
	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. using default Authenticator for now
		r.Use(jwtauth.Authenticator)

		
	})

	return r
}

