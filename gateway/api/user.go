package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	"github.com/ubclaunchpad/pinpoint/protobuf"
	"go.uber.org/zap"
)

type UserRouter struct {
	l *zap.SugaredLogger
	c pinpoint.CoreClient

	mux *chi.Mux
}

func newUserRouter(l *zap.SugaredLogger, c pinpoint.CoreClient) *UserRouter {
	router := chi.NewRouter()
	u := &UserRouter{l, c, router}
	router.Post("/create_user", u.createUser)
	// router.Get("/getUser/{userID}",getUser)
	return &UserRouter{l.Named("users"), c, router}
	// store email, name, and password, create a new schema package to store the user information model
}

func (u *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.mux.ServeHTTP(w, r)
}

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// parse request data
	var userData schema.User
	if err := decoder.Decode(&userData); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "invalid input"))
		return
	}
	// create user with data
	schema.NewUser(userData.Name, userData.Email, userData.Password)
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]string{
		"msg": "user created",
	})
}
