package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// API defines the API server. It is primarily a REST interface through which
// service.Service can be accessed.
type API struct {
	l *zap.SugaredLogger
	r *chi.Mux
	c pinpoint.CoreClient

	srv *http.Server
}

// New creates a new API server - start it using Run(). Returns a callback to
// close connection
func New(logger *zap.SugaredLogger) (*API, error) {
	router := chi.NewRouter()
	a := &API{
		l: logger.Named("api"),
		r: router,
		srv: &http.Server{
			Handler: router,
		},
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
	a.r.Get("/user/create_account", a.createAccountHandler)
	a.r.Get("/user/verify", a.verifyHandler)
}

// RunOpts defines options for API server startup
type RunOpts struct {
	GatewayOpts
	CoreOpts
}

// GatewayOpts defines gateway configuration options
type GatewayOpts struct {
	CertFile string
	KeyFile  string
}

// CoreOpts defines options for connecting to pinpoint-core
type CoreOpts struct {
	Host     string
	Port     string
	CertFile string
}

// Run spins up the API server
func (a *API) Run(host, port string, opts RunOpts) error {
	if host == "" && port == "" {
		return errors.New("invalid host and port configuration provided")
	}

	// set up server
	a.srv.Addr = host + ":" + port

	// set up parameters
	dialOpts := make([]grpc.DialOption, 0)
	if opts.CoreOpts.CertFile != "" {
		creds, err := credentials.NewClientTLSFromFile(opts.CoreOpts.CertFile, "")
		if err != nil {
			return fmt.Errorf("could not load tls cert: %s", err)
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	// connect to core service
	a.l.Infow("connecting to core",
		"core.host", opts.CoreOpts.Host,
		"core.port", opts.CoreOpts.Port,
		"core.tls", opts.CoreOpts.CertFile != "")
	conn, err := grpc.Dial(opts.Host+":"+opts.Port, dialOpts...)
	if err != nil {
		return fmt.Errorf("failed to connect to core service: %s", err.Error())
	}
	a.c = pinpoint.NewCoreClient(conn)
	defer conn.Close()

	// attempt connection
	go func() {
		if _, err = a.c.GetStatus(context.Background(), &request.Status{}); err != nil {
			a.l.Errorw("unable to connect to core service",
				"error", err.Error())
		} else {
			a.l.Info("established connection to core")
		}
	}()

	// lets gooooo
	tlsEnabled := opts.GatewayOpts.CertFile != ""
	a.l.Infow("spinning up api server",
		"gateway.host", host,
		"gateway.port", port,
		"gateway.tls", tlsEnabled)
	if tlsEnabled {
		if err := a.srv.ListenAndServeTLS(
			opts.GatewayOpts.CertFile, opts.GatewayOpts.KeyFile,
		); err != nil && err != http.ErrServerClosed {
			a.l.Warnw("error encountered - service stopped",
				"error", err)
			return err
		}
	} else {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.l.Warnw("error encountered - service stopped",
				"error", err)
			return err
		}
	}

	// report shutdown
	a.l.Info("service shut down")
	return nil
}

// Stop releases resources and shuts down the API server
func (a *API) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	if err := a.srv.Shutdown(ctx); err != nil {
		a.l.Warnw("error encountered during shutdown",
			"error", err.Error())
	}
	cancel()
}
