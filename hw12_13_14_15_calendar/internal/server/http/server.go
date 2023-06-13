package internalhttp

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/app"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	logger     logger.Logger
	config     configuration.ServerConfiguration
	httpServer *http.Server
	app        app.App
	validate   *validator.Validate
}

func NewServer(l logger.Logger, app app.App, cfg configuration.ServerConfiguration) *Server {
	v := validator.New()
	return &Server{logger: l, config: cfg, app: app, validate: v}
}

func (s *Server) Start(ctx context.Context) error {
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))

	s.httpServer = &http.Server{
		Addr:         address,
		Handler:      s.getServeMux(s.logger),
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
		IdleTimeout:  s.config.IdleTimeout,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// normal interrupt operation, ignored
				s.logger.Debug("http server stopped")
				return
			}
			s.logger.Fatalf("can not start http server: %+v", err)
		}
	}()

	s.logger.Debugf("http server started and listen %s", address)

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Debug("http server is shutting down...")
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) getServeMux(logg logger.Logger) *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", logRequest(http.HandlerFunc(notImplemented), logg))
	serveMux.Handle("/error", logRequest(http.HandlerFunc(serverError), logg))
	serveMux.Handle("/hello", logRequest(http.HandlerFunc(hello), logg))
	serveMux.Handle(app.URLPattern, logRequest(http.HandlerFunc(s.handleCalendarRequest), logg))
	return serveMux
}
