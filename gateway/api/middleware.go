package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// zapMiddleware manages logging requests and errors going through gin router
func zapMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	// use faster, default zap.Logger
	l := logger.Desugar().Named("gin")
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		var (
			// measure latency
			latency = time.Since(start)

			// gather metadata
			status = c.Writer.Status()
			method = c.Request.Method
			path   = c.Request.URL.Path
			query  = c.Request.URL.RawQuery
			ip     = c.ClientIP()
			agent  = c.Request.UserAgent()
		)

		if len(c.Errors) > 0 {
			// log as error with metadata if there are errors
			l.Error("error at "+path,
				zap.Strings("errors", c.Errors.Errors()),
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", ip),
				zap.String("user-agent", agent),
				zap.Duration("latency", latency))
		} else {
			// log as info with metadata by default
			l.Info("request at "+path,
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", ip),
				zap.String("user-agent", agent),
				zap.Duration("latency", latency))
		}
	}
}
