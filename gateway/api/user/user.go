package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
)

// Router routes to all user endpoints
type Router struct {
	l   *zap.SugaredLogger
	c   pinpoint.CoreClient
	mux *chi.Mux
}

// NewUserRouter instantiates a new router for handling user functionality
func NewUserRouter(l *zap.SugaredLogger, core pinpoint.CoreClient) *Router {
	u := &Router{l.Named("users"), core, chi.NewRouter()}

	// these should all be public
	u.mux.Post("/create", u.createUser)
	u.mux.Post("/login", u.login)
	u.mux.Get("/verify", u.verify)

	return u
}

func (u *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.mux.ServeHTTP(w, r)
}

func (u *Router) createUser(w http.ResponseWriter, r *http.Request) {
	// parse request data
	decoder := json.NewDecoder(r.Body)
	var user request.CreateAccount
	if err := decoder.Decode(&user); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}

	// create account in core
	resp, err := u.c.CreateAccount(r.Context(), &user)
	if err != nil {
		render.Render(w, r, res.ErrInternalServer("failed to create user account", err))
		return
	}

	// success!
	render.Render(w, r, res.Msg(resp.GetMessage(), http.StatusCreated,
		"email", user.GetEmail()))
}

type logininfo struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (u *Router) login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	var info logininfo
	json.Unmarshal([]byte(bodyString), &info)
	print(bodyString)

	email := info.Email
	password := info.Password

	// email := r.FormValue("email")
	// password := r.FormValue("password")

	if email == "" || password == "" {
		render.Render(w, r, res.ErrBadRequest("missing fields - both email and password is required"))
		return
	}

	if _, err := u.c.Login(r.Context(), &request.Login{
		Email: email, Password: password,
	}); err != nil {
		render.Render(w, r, res.ErrUnauthorized(err.Error()))
		return
	}

	// No error means authenticated, proceed to generate token
	// TODO: Generate token. See #10
	render.Render(w, r, res.MsgOK("user logged in",
		"token", 1234))
}

func (u *Router) verify(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")
	if hash == "" {
		render.Render(w, r, res.ErrBadRequest("hash is required"))
		return
	}

	resp, err := u.c.Verify(r.Context(), &request.Verify{Hash: hash})
	if err != nil {
		render.Render(w, r, res.ErrNotFound(err.Error()))
		return
	}

	render.Render(w, r, res.Msg(resp.GetMessage(), http.StatusAccepted))
}
