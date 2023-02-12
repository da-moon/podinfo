package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/armon/go-metrics"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

// Metrics tracks how long it takes for a request to complete
func Metrics(pattern string, l *logger.WrappedLogger) func(next http.HandlerFunc) http.HandlerFunc {
	pathLabel := strings.Replace(pattern[1:], "/", "_", -1)
	l.Info("metrics middleware path = [ %s ]  label = [ %s ]", pattern, pathLabel)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			labels := []metrics.Label{{Name: "method", Value: r.Method}, {Name: "path", Value: pathLabel}}
			metrics.MeasureSinceWithLabels([]string{"api", "http"}, start, labels)
		})
	}
}
