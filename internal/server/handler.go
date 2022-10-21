package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"token-repository/internal/model"
	"token-repository/internal/usecase"

	"github.com/gorilla/mux"
)

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	pathParam := mux.Vars(r)
	tokenName := pathParam["token_name"]
	token, err := s.Service.GetToken(tokenName)
	if err != nil {
		if errors.Is(err, usecase.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "token not found\n")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	outputJson, err := json.Marshal(&token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outputJson))
}

func (s *Server) updateHandler(w http.ResponseWriter, r *http.Request) {
	var token model.OAuth2Update
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = s.Service.UpdateToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "token recorded\n")
}
