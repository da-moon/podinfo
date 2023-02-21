package middlewares

import (
	"bytes"
	"fmt"
	"net/http"

	logger "github.com/da-moon/podinfo/internal/logger"
)

// Log uses wrapped logger for logging requests
func Log(l *logger.WrappedLogger) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var logString bytes.Buffer
			if r.Host != "" {
				fmt.Fprintf(&logString, "host=%s", r.Host)
			}
			if r.RemoteAddr != "" {
				fmt.Fprintf(&logString, " address=%s", r.RemoteAddr)
			}
			if r.Method != "" {
				fmt.Fprintf(&logString, " method=%s", r.Method)
			}
			if r.RequestURI != "" {
				fmt.Fprintf(&logString, " request_uri=%s", r.RequestURI)
			}
			if r.Proto != "" {
				fmt.Fprintf(&logString, " proto=%s", r.Proto)
			}
			if r.UserAgent() != "" {
				fmt.Fprintf(&logString, " user_agent=%s", r.Proto)
			}
			l.Info(logString.String())
			next.ServeHTTP(w, r)
		})
	}
}
