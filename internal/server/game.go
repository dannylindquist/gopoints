package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) registerGameRoutes() {
	s.router.HandleFunc("/game/{id}", s.gameById).Methods("GET")
	s.router.HandleFunc("/game", s.createGame).Methods("POST")
}

func (s *Server) createGame(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var content map[string]interface{}
	json.Unmarshal(body, &content)
	fmt.Printf("%v\n", content)
	w.Header().Add("cOntent-type", "application/json")
	w.Write(body)
}

func (s *Server) gameById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Fprintf(w, "bad request: %v", err)
	}
	game := s.GameService.GetGame(r.Context(), id)
	if game == nil {
		http.Error(w, "game not found", http.StatusNotFound)
		return
	} else {
		content, err := json.Marshal(game)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(content)
	}
}
