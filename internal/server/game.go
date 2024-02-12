package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dannylindquist/gopoints"
	"github.com/dannylindquist/gopoints/internal/server/views"
	"github.com/gorilla/mux"
)

func (s *Server) registerGameRoutes() {
	s.router.HandleFunc("/", s.home).Methods("GET")
	s.router.HandleFunc("/game/{id}", s.gameById).Methods("GET")
	s.router.HandleFunc("/game", s.createGame).Methods("POST")
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "main.html", nil)
}

func (s *Server) createGame(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	game := s.GameService.CreateGame(r.Context(), gopoints.GameUpdate{Name: name, Type: "ASC"})
	w.Header().Add("Hx-Location", fmt.Sprintf("/game/%d", game.ID))
}

func (s *Server) gameById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Fprintf(w, "bad request: %v", err)
		return
	}
	game := s.GameService.GetGame(r.Context(), id)
	if game == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	w.Header().Add("content-type", "text/html")
	views.Render(w, "game.html", game)
}
