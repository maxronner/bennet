package rest

import "github.com/maxronner/bennet/internal/app/responder"

func NewServer(service *responder.Service) *Server {
	return &Server{
		service,
	}
}

type Server struct {
	service *responder.Service
}