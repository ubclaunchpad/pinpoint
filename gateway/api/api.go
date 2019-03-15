package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ubclaunchpad/pinpoint/gateway/api/club"
	"github.com/ubclaunchpad/pinpoint/gateway/api/user"
	"github.com/ubclaunchpad/pinpoint/gateway/utils"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// API defines the API server. It is primarily a REST interface through which
// service.Service can be accessed.
type API struct {
	l *zap.SugaredLogger
	r *chi.Mux
	c pinpoint.CoreClient

	srv  *http.Server
	grpc *grpc.ClientConn
}

// CoreOpts defines options for connecting to pinpoint-core
type CoreOpts struct {
	Host     string
	Port     string
	CertFile string
	Token    string
}

// New creates a new API server - start it using Run(). Returns a callback to
// close connection
func New(logger *zap.SugaredLogger, opts CoreOpts) (*API, error) {
	router := chi.NewRouter()
	a := &API{
		l: logger.Named("api"),
		r: router,
		srv: &http.Server{
			Handler: router,
		},
	}

	// set up core client
	if err := a.setUpCoreClient(opts); err != nil {
		return nil, err
	}

	// set up endpoints
	a.setUpRouter()
	a.registerHandlers()

	return a, nil
}

// setUpCoreClient initializes the API server's clients
func (a *API) setUpCoreClient(opts CoreOpts) error {
	// set up parameters for core conn
	dialOpts := make([]grpc.DialOption, 0)
	if opts.CertFile != "" {
		creds, err := credentials.NewClientTLSFromFile(opts.CertFile, "")
		if err != nil {
			return fmt.Errorf("could not load tls cert: %s", err)
		}
		dialOpts = append(dialOpts,
			grpc.WithTransportCredentials(creds),
			grpc.WithPerRPCCredentials(utils.NewCredentials(opts.Token)))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	// connect to core service
	a.l.Infow("setting up core client",
		"core.host", opts.Host,
		"core.port", opts.Port,
		"core.tls", opts.CertFile != "")
	var err error
	a.grpc, err = grpc.Dial(opts.Host+":"+opts.Port, dialOpts...)
	if err != nil {
		return fmt.Errorf("failed to connect to core service: %s", err.Error())
	}
	a.c = pinpoint.NewCoreClient(a.grpc)
	return nil
}

// setUpRouter initializes any middleware or general things the API router
// might need
func (a *API) setUpRouter() {
	// CORS setting for development use
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	a.r.Use(
		c.Handler,
		middleware.RequestID,
		middleware.RealIP,
		newLoggerMiddleware("router", a.l),
		middleware.Recoverer)
}

// registerHandler sets up server routes
func (a *API) registerHandlers() {
	a.r.Get("/status", a.statusHandler)
	a.r.Mount("/user", user.NewUserRouter(a.l, a.c))
	a.r.Mount("/club", club.NewClubRouter(a.l, a.c))
}

// runs Core and Gateway connection handshake
func (a *API) establishConnection(ctx context.Context, token string) error {
	var header metadata.MD
	if _, err := a.c.Handshake(ctx, &request.Empty{},
		grpc.Header(&header)); err != nil {
		return fmt.Errorf("error during handshake: %s", err.Error())
	}
	if tokens := header.Get("authorization"); tokens == nil || len(tokens) == 0 || tokens[0] != token {
		return errors.New("core provided invalid authentication")
	}
	return nil
}

// RunOpts defines options for API server startup
type RunOpts struct {
	CertFile string
	KeyFile  string
	Token    string
}

// Run spins up the API server
func (a *API) Run(host, port string, opts RunOpts) error {
	if host == "" && port == "" {
		return errors.New("invalid host and port configuration provided")
	}

	// set up server
	a.srv.Addr = host + ":" + port

	// initial validation
	go func() {
		if err := a.establishConnection(context.Background(), opts.Token); err != nil {
			a.l.Errorw("unable to connect to core service",
				"error", err.Error())
		} else {
			a.l.Info("successfully connected to core service")
		}
	}()

	// lets gooooo
	tlsEnabled := opts.CertFile != ""
	a.l.Infow("spinning up api server",
		"gateway.host", host,
		"gateway.port", port,
		"gateway.tls", tlsEnabled)
	if tlsEnabled {
		if err := a.srv.ListenAndServeTLS(
			opts.CertFile, opts.KeyFile,
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
	if a.grpc != nil {
		a.grpc.Close()
	}

	// hack to make sure everything stops properly, notably API::Run()
	time.Sleep(time.Second)
}
