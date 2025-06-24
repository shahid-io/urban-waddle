package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shahid-io/urban-waddle/service/user"
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

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler(s.db)
	userHandler.RegisterRoutes(subrouter)

	log.Println("API server running on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}
