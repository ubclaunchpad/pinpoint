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

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// parse request data
	var userData schema.User
	err := decoder.Decode(&userData)
	if err != nil {
		panic(err)
	}
	// create user with data
	schema.NewUser(userData.Name, userData.Email, userData.Password)

	resp, err := u.c.GetStatus(context.Background(), &request.Status{Callback: "User Created"})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)

}
