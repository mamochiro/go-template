package grpcserver

import (
	"fmt"
	"go-boilerplate/internals/config"
	"go-boilerplate/internals/controller"
	"net"
	"strconv"

	api_v1 "go-boilerplate/pkg/api/v1"
	grpc_health_v1 "go-boilerplate/pkg/grpc/health/v1"

	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	Config       config.Configuration
	Server       *grpc.Server
	HealthCtrl   *controller.HealthZController
	PingPongCtrl *controller.PingPongController
}

// Configure ...
func (s *Server) Configure() {
	grpc_health_v1.RegisterHealthServer(s.Server, s.HealthCtrl)
	api_v1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
}

// Start ...
func (s *Server) Start() {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(s.Config.Port))
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening and serving HTTP on", strconv.Itoa(s.Config.Port))
	if err := s.Server.Serve(listen); err != nil {
		panic(err)
	}
}

// NewServer ...
func NewServer(config config.Configuration, healthCtrl *controller.HealthZController, pingPongCtrl *controller.PingPongController) *Server {
	s := &Server{
		Server:       grpc.NewServer(),
		Config:       config,
		HealthCtrl:   healthCtrl,
		PingPongCtrl: pingPongCtrl,
	}
	s.Configure()
	return s
}
