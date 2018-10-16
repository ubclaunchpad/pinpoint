package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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
	_, err = a.c.GetStatus(context.Background(), &request.Status{})
	if err != nil {
		return fmt.Errorf("failed to connect to core service: %s", err.Error())
	}

	// lets gooooo
	a.l.Infow("spinning up api server",
		"gateway.host", host,
		"gateway.port", port,
		"gateway.tls", opts.GatewayOpts.CertFile != "")
	addr := host + ":" + port
	if opts.GatewayOpts.CertFile != "" {
		err = http.ListenAndServeTLS(
			addr,
			opts.GatewayOpts.CertFile,
			opts.GatewayOpts.KeyFile,
			a.r)
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
