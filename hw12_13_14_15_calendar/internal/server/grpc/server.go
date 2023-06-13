//go:generate protoc EventService.proto --proto_path=./../../../api --go_out=.  --go-grpc_out=.

package internalgrpc

import (
	"context"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/app"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	UnimplementedEvensServer
	logger     logger.Logger
	app        app.App
	grpcServer *grpc.Server
	config     configuration.GrpcConfiguration
}

func NewServer(logger logger.Logger, app app.App, configuration configuration.GrpcConfiguration) *Server {
	server := &Server{logger: logger, app: app, config: configuration}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(server.unaryInterceptor),
		grpc.KeepaliveEnforcementPolicy(configuration.EnforcementPolicy),
		grpc.KeepaliveParams(configuration.ServerParameters))
	server.grpcServer = grpcServer
	RegisterEvensServer(server.grpcServer, server)
	return server
}

func (s Server) Start(ctx context.Context) error {
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			s.logger.Fatalf("can not start http server: %+v", err)
		}
		s.logger.Debug("grpc server stopped")
	}()

	s.logger.Debugf("grpc server started and listen %s", address)
	<-ctx.Done()

	return nil
}

func (s Server) Stop() {
	s.logger.Debug("grpc server is shutting down...")
	s.grpcServer.GracefulStop()
}

func (s Server) CreateEvent(_ context.Context, request *NewEventRequest) (*EventIdResponse, error) {
	userID := uuid.New()
	err := userID.UnmarshalBinary(request.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	event := &storage.Event{
		Title:    request.Title,
		Time:     request.Time.AsTime(),
		Duration: time.Duration(request.Duration),
		UserID:   userID,
	}

	eventID, err := s.app.CreateEvent(event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	binary, err := eventID.MarshalBinary()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &EventIdResponse{Id: binary}, err
}

func (s Server) UpdateEvent(_ context.Context, request *ChangeEventRequest) (*empty.Empty, error) {
	eventID := uuid.New()
	err := eventID.UnmarshalBinary(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	userID := uuid.New()
	err = userID.UnmarshalBinary(request.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	event := &storage.Event{
		Title:    request.Title,
		Time:     request.Time.AsTime(),
		Duration: time.Duration(request.Duration),
		UserID:   userID,
	}

	err = s.app.UpdateEvent(eventID, event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s Server) DeleteEvent(_ context.Context, request *EventIdRequest) (*empty.Empty, error) {
	eventID := uuid.New()
	err := eventID.UnmarshalBinary(request.Id)
	if errors.Is(err, storage.ErrEventNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.app.DeleteEvent(eventID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s Server) DayEvens(_ context.Context, request *TimeRequest) (*EventsResponse, error) {
	appEvents, err := s.app.DayEvens(request.Time.AsTime())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	events, err := s.mapEvents(appEvents)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s Server) WeekEvens(_ context.Context, request *TimeRequest) (*EventsResponse, error) {
	appEvents, err := s.app.WeekEvens(request.Time.AsTime())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	events, err := s.mapEvents(appEvents)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s Server) MonthEvens(_ context.Context, request *TimeRequest) (*EventsResponse, error) {
	appEvents, err := s.app.MonthEvens(request.Time.AsTime())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	events, err := s.mapEvents(appEvents)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s Server) mapEvents(events []*storage.Event) (*EventsResponse, error) {
	response := &EventsResponse{}
	response.Events = make([]*EventResponse, 0, len(events))
	for _, event := range events {
		eventID, err := event.ID.MarshalBinary()
		if err != nil {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		userID, err := event.UserID.MarshalBinary()
		if err != nil {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		e := &EventResponse{
			Id:       eventID,
			Title:    event.Title,
			Time:     timestamppb.New(event.Time),
			Duration: event.Duration.Milliseconds() / 1000,
			UserId:   userID,
		}
		response.Events = append(response.Events, e)
	}
	return response, nil
}

func (s Server) unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	m, err := handler(ctx, req)
	if err != nil {
		s.logger.Errorf("RPC '%s' failed with error: %+v", info.FullMethod, err)
		return m, err
	}
	s.logger.Debugf("RPC '%s', %s", info.FullMethod, time.Since(start))
	return m, err
}
