package api


import (
	"net/http"
	"github.com/ubclaunchpad/pinpoint/gateway/auth"
)

// LoginHandler login user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// how to implement this login handler ?
	// assuming user logged in, then I should generate token

	//
	token, _ := auth.GenerateToken();

	// then how to return it to client?

}
