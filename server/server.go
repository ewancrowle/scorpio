package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scorpio/routes"
	"strconv"
	"time"
)

type Server struct {
	router *mux.Router
	server *http.Server
}

func (s Server) Run(port int) error {
	s.router = NewRouter()
	s.server = &http.Server{
		Handler:        s.router,
		Addr:           "127.0.0.1:" + strconv.Itoa(port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Server initalised with port: %d\n", port)
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
