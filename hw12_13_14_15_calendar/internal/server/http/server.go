package internalhttp

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	logger     logger.Logger
	config     configuration.ServerConfiguration
	httpServer *http.Server
}

type Application interface { // TODO
}

func NewServer(l logger.Logger, app Application, cfg configuration.ServerConfiguration) *Server { //nolint:revive
	return &Server{logger: l, config: cfg}
}

func (s *Server) Start(ctx context.Context) error {
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))

	s.httpServer = &http.Server{
		Addr:         address,
		Handler:      getServeMux(s.logger),
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
		IdleTimeout:  s.config.IdleTimeout,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// normal interrupt operation, ignored
				s.logger.Debug("server stopped")
				return
			}
			s.logger.Fatalf("can not start http server: %+v", err)
		}
	}()

	s.logger.Debugf("server started and listen http://%s", address)

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Debug("server is shutting down...")
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func getServeMux(logg logger.Logger) *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", logRequest(http.HandlerFunc(NotImplemented), logg))
	serveMux.Handle("/hello", logRequest(http.HandlerFunc(Hello), logg))
	return serveMux
}
