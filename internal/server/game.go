package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) registerGameRoutes() {
	s.router.HandleFunc("/game/{id}", s.gameById)
}

func (s *Server) gameById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Fprintf(w, "bad request: %v", err)
	}
	fmt.Fprint(w, id)
}
