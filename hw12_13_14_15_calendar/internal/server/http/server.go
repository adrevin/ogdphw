package internalhttp

import (
	"context"
	"fmt"
	"net"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
)

type Server struct { // TODO
}

type Application interface { // TODO
}

var (
	logg   logger.Logger
	config configuration.ServerConfiguration
)

func NewServer(l logger.Logger, app Application, cfg configuration.ServerConfiguration) *Server { //nolint:revive
	logg = l
	config = cfg
	logg.Debug("server created")
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	address := net.JoinHostPort(config.Host, config.Port)
	logg.Debug(fmt.Sprintf("server started and lisen %s", address))

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error { //nolint:revive
	// TODO
	return nil
}

// TODO
