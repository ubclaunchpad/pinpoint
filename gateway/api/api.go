package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	pinpoint "github.com/ubclaunchpad/pinpoint/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// API defines the API server. It is primarily a REST interface through which
// service.Service can be accessed.
type API struct {
	l *zap.SugaredLogger
	r *chi.Mux
	c pinpoint.PinpointCoreClient
}

// New creates a new API server - start it using Run()
func New(conn *grpc.ClientConn, logger *zap.SugaredLogger, debug bool) (*API, error) {
	a := &API{
		r: chi.NewRouter(), l: logger.Named("api"),
		c: pinpoint.NewPinpointCoreClient(conn),
	}

	a.setUpRouter()
	a.registerHandlers()

	return a, nil
}

// setUpRouter initializes any middleware or general things the API router
// might need
func (a *API) setUpRouter() {
	a.r.Use(
		middleware.RequestID,
		middleware.RealIP,
		newLoggerMiddleware("router", a.l),
		middleware.Recoverer)
}

// registerHandler sets up server routes
func (a *API) registerHandlers() {
	a.r.Get("/status", a.statusHandler)
}

// RunOpts defines options for API server startup
type RunOpts struct {
	*SSLOpts
	*ServiceOpts
}

// SSLOpts defines SSL options
type SSLOpts struct {
	CertFile string
	KeyFile  string
}

// ServiceOpts defines options for connecting to the pinpoint service
type ServiceOpts struct {
}

// Run spins up the API server
func (a *API) Run(host, port string, opts RunOpts) error {
	if host == "" && port == "" {
		return errors.New("invalid host and port configuration provided")
	}

	a.l.Infow("spinning up api server",
		"tls", opts.SSLOpts != nil,
		"host", host,
		"port", port)

	addr := host + ":" + port
	if opts.SSLOpts != nil && opts.CertFile != "" {
		return http.ListenAndServeTLS(addr, opts.CertFile, opts.KeyFile, a.r)
	}
	return http.ListenAndServe(addr, a.r)
}
