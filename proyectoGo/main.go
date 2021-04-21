package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"adminModPerl/database"
	"adminModPerl/routes"
)

type Server struct {
	Router *mux.Router
	Addr   string
}

func (s *Server) Initialize(addr string) {
	s.Router = routes.RegisterRoutes()
	s.Addr = addr
}

func (s *Server) Run() {
	log.Println("Server running on", s.Addr)
	http.Handle("/", s.Router)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func main() {
	database.Migrate()
	if len(database.GetRoles()) <= 0 {
		database.Inicializar()
	}
	server := Server{}
	server.Initialize(":5051")
	server.Run()
}
