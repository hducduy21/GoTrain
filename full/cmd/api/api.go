package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hducduy21/gofull/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, subrouter)
}
