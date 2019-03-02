package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/api/ctxutil"
	"github.com/ubclaunchpad/pinpoint/gateway/auth"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

var tokenAuth *jwtauth.JWTAuth

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

	// Authenticated endpoints
	u.mux.Group(func(r chi.Router) {
		// JWT Initialization
		key, err := auth.GetAPIPrivateKey()
		if err != nil {
			log.Fatal(err.Error())
		}
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(jwtauth.New("HS256", key, nil)))
		// Handle valid/invalid tokens
		r.Use(jwtauth.Authenticator)
		r.Post("/verify", u.verify)
	})

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
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}

	// create account in core
	resp, err := u.c.CreateAccount(r.Context(), &user)
	if err != nil {
		l.Debugw("error occured creating user account", "error", err)
		st, ok := status.FromError(err)
		if !ok {
			render.Render(w, r, res.ErrInternalServer("failed to create user account", err))
			return
		}

		switch st.Code() {
		case codes.InvalidArgument:
			render.Render(w, r, res.ErrBadRequest(st.Message()))
		default:
			render.Render(w, r, res.ErrInternalServer(st.Message(), err))
		}
		return
	}

	// success!
	render.Render(w, r, res.Msg(resp.GetMessage(), http.StatusCreated,
		"email", user.GetEmail()))
}

func (u *Router) login(w http.ResponseWriter, r *http.Request) {
	var l = u.l.With("request-id", ctxutil.GetRequestID(r))
	if r.Body == nil {
		render.Render(w, r, res.ErrBadRequest("missing request body"))
		return
	}

	var decoder = json.NewDecoder(r.Body)
	var info struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := decoder.Decode(&info); err != nil {
		render.Render(w, r, res.ErrBadRequest("error occurred parsing user login form entry",
			"error", err))
		return
	}

	if info.Email == "" || info.Password == "" {
		render.Render(w, r, res.ErrBadRequest("missing fields - both email and password is required"))
		return
	}

	if _, err := u.c.Login(r.Context(), &request.Login{
		Email: info.Email, Password: info.Password,
	}); err != nil {
		render.Render(w, r, res.ErrUnauthorized(err.Error()))
		return
	}

	// No error means authenticated, proceed to generate token
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &auth.Claims{
		Email: info.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenStr, err := claims.GenerateToken()
	if err != nil {
		render.Render(w, r, res.ErrInternalServer("failed to generate token", err))
		return
	}

	l.Infow("logged in", "user", info.Email)
	render.Render(w, r, res.MsgOK("user logged in",
		"token", tokenStr))
}

func (u *Router) verify(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	email := fmt.Sprintf("%v", claims["email"])
	hash := r.FormValue("hash")
	if hash == "" {
		render.Render(w, r, res.ErrBadRequest("hash is required"))
		return
	}
	resp, err := u.c.Verify(r.Context(), &request.Verify{Email: email, Hash: hash})
	if err != nil {
		render.Render(w, r, res.ErrNotFound(err.Error()))
		return
	}

	render.Render(w, r, res.Msg(resp.GetMessage(), http.StatusAccepted))
}
