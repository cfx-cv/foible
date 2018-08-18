package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	nats "github.com/nats-io/go-nats"
)

type Server struct {
	conn *nats.Conn
	logs []string
}

func NewServer(conn *nats.Conn) *Server {
	return &Server{conn: conn}
}

func (s *Server) Start() {
	s.subscribe()

	router := mux.NewRouter()
	router.HandleFunc("/", s.index).Methods("GET")

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal(err)
	}
}
