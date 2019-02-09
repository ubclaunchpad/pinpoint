package user

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/api/ctxutil"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
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
	var l = u.l.With("request-id", ctxutil.GetRequestID(r))
	// parse request data
	decoder := json.NewDecoder(r.Body)
	var user request.CreateAccount
	if err := decoder.Decode(&user); err != nil {
		l.Debugw("error occured reading request", "error", err)
		render.Render(w, r, res.ErrBadRequest(r, "invalid request"))
		return
	}

	// create account in core
	resp, err := u.c.CreateAccount(r.Context(), &user)
	if err != nil {
		l.Debugw("error occured creating user account", "error", err)
		st, ok := status.FromError(err)
		if !ok {
			render.Render(w, r, res.ErrInternalServer(r, "failed to create user account",
				"error", err.Error()))
			return
		}

		switch st.Code() {
		case codes.InvalidArgument:
			render.Render(w, r, res.ErrBadRequest(r, st.Message()))
		default:
			render.Render(w, r, res.ErrInternalServer(r, st.Message(),
				"error", err.Error()))
		}
		return
	}

	// success!
	render.Render(w, r, res.Message(r, resp.GetMessage(), http.StatusCreated,
		"email", user.GetEmail()))
}

func (u *Router) login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		render.Render(w, r, res.ErrBadRequest(r, "missing fields - both email and password is required"))
		return
	}

	if _, err := u.c.Login(r.Context(), &request.Login{
		Email: email, Password: password,
	}); err != nil {
		render.Render(w, r, res.ErrUnauthorized(r, err.Error()))
		return
	}

	// No error means authenticated, proceed to generate token
	w.WriteHeader(http.StatusOK)
	// TODO: Generate token. See #10
	render.JSON(w, r, map[string]string{
		"token": "1234",
	})
}

func (u *Router) verify(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")
	if hash == "" {
		render.Render(w, r, res.ErrBadRequest(r, "hash is required"))
		return
	}

	resp, err := u.c.Verify(r.Context(), &request.Verify{Hash: hash})
	if err != nil {
		render.Render(w, r, res.Err(r, err.Error(), http.StatusNotFound))
		return
	}

	render.Render(w, r, res.Message(r, resp.GetMessage(), http.StatusAccepted))
}
