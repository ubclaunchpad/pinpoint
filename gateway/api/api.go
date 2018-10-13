package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// API defines the API server. It is primarily a REST interface through which
// service.Service can be accessed.
type API struct {
	l *zap.SugaredLogger
	r *chi.Mux
	c pinpoint.CoreClient
}

// New creates a new API server - start it using Run(). Returns a callback to
// close connection
func New(logger *zap.SugaredLogger) (*API, error) {
	a := &API{
		r: chi.NewRouter(), l: logger.Named("api"),
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
	SSLOpts
	CoreOpts
}

// SSLOpts defines SSL options
type SSLOpts struct {
	CertFile string
	KeyFile  string
}

// CoreOpts defines options for connecting to pinpoint-core
type CoreOpts struct {
	Host        string
	Port        string
	DialOptions []grpc.DialOption
}

// Run spins up the API server
func (a *API) Run(host, port string, opts RunOpts) error {
	if host == "" && port == "" {
		return errors.New("invalid host and port configuration provided")
	}

	// connect to core server
	a.l.Infow("connecting to core",
		"core.host", opts.Host,
		"core.port", opts.Port)
	conn, err := grpc.Dial(opts.Host+":"+opts.Port, opts.DialOptions...)
	if err != nil {
		return fmt.Errorf("failed to connect to core service: %s", err.Error())
	}
	a.c = pinpoint.NewCoreClient(conn)
	defer conn.Close()

	// lets gooooo
	a.l.Infow("spinning up api server",
		"tls", opts.CertFile != "",
		"host", host,
		"port", port)
	addr := host + ":" + port
	if opts.CertFile != "" {
		err = http.ListenAndServeTLS(addr, opts.CertFile, opts.KeyFile, a.r)
	} else {
		err = http.ListenAndServe(addr, a.r)
	}
	if err != nil {
		a.l.Errorf("error encountered - service stopped",
			"error", err)
		return err
	}

	// report shutdown
	a.l.Info("service shut down")
	return nil
}
