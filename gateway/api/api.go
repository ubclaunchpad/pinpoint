package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	pinpoint "github.com/ubclaunchpad/pinpoint/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// API defines the API server. It is primarily a REST interface through which
// service.Service can be accessed.
type API struct {
	l *zap.SugaredLogger
	r *gin.Engine
	c pinpoint.PinpointCoreClient
}

// New creates a new API server - start it using Run()
func New(conn *grpc.ClientConn, logger *zap.SugaredLogger, debug bool) (*API, error) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	a := &API{
		r: gin.New(), l: logger.Named("api"),
		c: pinpoint.NewPinpointCoreClient(conn),
	}

	a.setUpEngine()
	a.registerHandlers()

	return a, nil
}

func (a *API) setUpEngine() {
	a.r.Use(zapMiddleware(a.l), gin.Recovery())
}

func (a *API) registerHandlers() {
	a.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
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
	if opts.SSLOpts != nil {
		return a.r.RunTLS(addr, opts.CertFile, opts.KeyFile)
	}
	return a.r.Run(addr)
}
