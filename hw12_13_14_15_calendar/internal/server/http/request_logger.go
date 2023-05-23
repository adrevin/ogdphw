package internalhttp

import (
	"net"
	"net/http"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	length     int
	logger     logger.Logger
}

func newResponseWriter(w http.ResponseWriter, logger logger.Logger) *responseWriter {
	return &responseWriter{w, http.StatusOK, 0, logger}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(bytes []byte) (int, error) {
	l, err := rw.ResponseWriter.Write(bytes)
	if err != nil {
		rw.logger.Errorf("write response error: %+v", err)
		return 0, err
	}
	rw.length += l
	return l, nil
}

func logRequest(next http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w, logger)
		next.ServeHTTP(rw, r)

		var address string
		address, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			address = "unknown"
		}
		logger.Debugf(
			"%s [%s] %s %s %s %d %d %s %s",
			address,
			time.Now().UTC(),
			r.Method,
			r.URL,
			r.Proto,
			rw.statusCode,
			rw.length,
			time.Since(start),
			r.UserAgent())
	})
}
