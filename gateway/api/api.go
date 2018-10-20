package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"google.golang.org/grpc/metadata"
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

// Runs Core and Gateway Connection Handshake
func (a *API) establishConnection() error {
	md := metadata.Pairs("token", os.Getenv("PINPOINT_CORE_TOKEN"))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	//Authenticate with core first
	var header, trailer metadata.MD
	var authflag bool
	_, err := a.c.HandShake(ctx, &request.Empty{}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		a.l.Errorf("Error when setting up handshake: %s", err)
	}
	for _, value := range header {
		if value[0] == os.Getenv("PINPOINT_GATEWAY_TOKEN") {
			a.l.Info("Core passed authentication")
			authflag = true
		}
	}
	if authflag != true {
		a.l.Info("Core failed authentication, connection closing")
		err = errors.New("Core failed authentication, connection closing")
	}

	return err
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

	// Exchange auth tokens with core
	if err := a.establishConnection(); err != nil {
		conn.Close()
	}

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
