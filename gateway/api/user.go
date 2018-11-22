package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
)

// UserRouter routes to all user endpoints
type UserRouter struct {
	l   *zap.SugaredLogger
	c   pinpoint.CoreClient
	mux *chi.Mux
}

func newUserRouter(l *zap.SugaredLogger, core pinpoint.CoreClient) *UserRouter {
	u := &UserRouter{l.Named("users"), core, chi.NewRouter()}

	// these should all be public
	u.mux.Post("/create", u.createUser)
	u.mux.Post("/login", u.login)
	u.mux.Get("/verify", u.verify)

	return u
}

func (u *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.mux.ServeHTTP(w, r)
}

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	// parse request data
	decoder := json.NewDecoder(r.Body)
	var user request.CreateAccount
	if err := decoder.Decode(&user); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "invalid request"))
		return
	}

	// create account in core
	resp, err := u.c.CreateAccount(context.Background(), &user)
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}

	// success!
	render.Render(w, r, res.Message(r, resp.GetMessage(), http.StatusCreated,
		"email", user.GetEmail()))
}

func (u *UserRouter) login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		render.Render(w, r, res.ErrBadRequest(r, errors.New("missing fields"), ""))
		return
	}

	resp, err := u.c.Login(context.Background(), &request.Login{Email: email, Password: password})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	if resp.GetMessage() != "user successfully logged in" {
		render.Render(w, r, res.Message(r, resp.GetMessage(), http.StatusUnauthorized))
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO: Generate token
	render.JSON(w, r, map[string]string{
		"token": "1234",
	})
}

func (u *UserRouter) verify(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")
	if hash == "" {
		render.Render(w, r, res.ErrBadRequest(r, errors.New("missing fields"), ""))
		return
	}

	resp, err := u.c.Verify(context.Background(), &request.Verify{Hash: hash})
	if err != nil {
		render.Render(w, r, res.Err(r, err, http.StatusNotFound))
		return
	}
	render.Render(w, r, res.Message(r, resp.GetMessage(), http.StatusAccepted))
}
