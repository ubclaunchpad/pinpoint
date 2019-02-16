package api

import (
	"time"

	"github.com/ubclaunchpad/pinpoint/gateway/api/ctxutil"

	"net/http"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type loggerMiddleware struct {
	l *zap.Logger
}

func newLoggerMiddleware(name string, logger *zap.SugaredLogger) func(next http.Handler) http.Handler {
	// use faster, default zap.Logger
	return loggerMiddleware{logger.Desugar().Named(name)}.Handler
}

// zapMiddleware manages logging requests and errors going through gin router
func (z loggerMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		latency := time.Since(start)

		z.l.Info("request completed",
			// request metadata
			zap.String("path", r.URL.Path),
			zap.String("query", r.URL.RawQuery),
			zap.String("method", r.Method),
			zap.String("user-agent", r.UserAgent()),

			// response metadata
			zap.Int("status", ww.Status()),
			zap.Duration("took", latency),

			// additional metadata
			zap.String("real-ip", r.RemoteAddr),
			zap.String("request-id", ctxutil.GetRequestID(r)))
	})
}
