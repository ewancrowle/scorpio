package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scorpio/routes"
	"time"
)

type Server struct {
	router *mux.Router
	server *http.Server
}

func (s Server) Run(address string) error {
	s.router = NewRouter()
	s.server = &http.Server{
		Addr:           address,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Server initalised on address: %s\n", address)
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.GetBlockchain).Methods("GET")
	return router
}
