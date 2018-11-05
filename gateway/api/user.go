package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	"github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
)

// UserRouter routes too all user endpoints
type UserRouter struct {
	l *zap.SugaredLogger
	c pinpoint.CoreClient

	mux *chi.Mux
}

func newUserRouter(l *zap.SugaredLogger, c pinpoint.CoreClient) *UserRouter {
	router := chi.NewRouter()
	u := &UserRouter{l, c, router}
	router.Post("/create_user", u.createUser)
	router.Post("/verify", u.verify)
	return &UserRouter{l.Named("users"), c, router}
}

func (u *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.mux.ServeHTTP(w, r)
}

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// parse request data
	var userData schema.CreateUser
	if err := decoder.Decode(&userData); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}
	resp, err := u.c.CreateAccount(context.Background(), &request.CreateAccount{
		Email:           userData.Email,
		Name:            userData.Name,
		Password:        userData.Password,
		ConfirmPassword: userData.CPassword,
		EmailSubscribe:  userData.ESub,
	})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	rJSON, err := json.Marshal(resp)
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
	}

	render.Render(w, r, res.Message(r, string(rJSON), http.StatusCreated))
}

func (u *UserRouter) verify(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")
	resp, err := u.c.Verify(context.Background(), &request.Verify{Hash: hash})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)
}
