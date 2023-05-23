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
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK, 0}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(bytes []byte) (int, error) {
	l, e := rw.ResponseWriter.Write(bytes)
	rw.length += l
	return l, e
}

func logRequest(next http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		remoteAddr := r.RemoteAddr
		header := r.Header.Get("X-Forwarded-For")
		realIP := net.ParseIP(header)
		if realIP != nil {
			remoteAddr = realIP.String()
		}

		logger.Debugf(
			"%s '%s %s', %d, %d bytes, %s",
			remoteAddr,
			r.Method,
			r.URL,
			rw.statusCode,
			rw.length,
			time.Since(start))
	})
}
