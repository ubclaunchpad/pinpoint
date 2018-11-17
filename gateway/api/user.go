package api

import (
	"context"
	"encoding/json"
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

func newUserRouter(l *zap.SugaredLogger, c pinpoint.CoreClient) *UserRouter {
	router := chi.NewRouter()
	u := &UserRouter{l, c, router}
	router.Post("/create", u.createUser)
	router.Post("/login", u.login)
	router.Get("/verify", u.verify)
	return &UserRouter{l.Named("users"), c, router}
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
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{
		"token": "1234",
	})
}

func (u *UserRouter) verify(w http.ResponseWriter, r *http.Request) {
	resp, err := u.c.Verify(context.Background(), &request.Verify{Hash: r.FormValue("hash")})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.Render(w, r, res.Message(r, resp.GetMessage(), http.StatusAccepted))
}
