package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()
	s := &Server{
		router,
	}

	s.registerGameRoutes()

	return s
}

func (s *Server) Open() {
	go http.ListenAndServe(":3000", s.router)
	fmt.Println("starting on http://localhost:3000")
}
