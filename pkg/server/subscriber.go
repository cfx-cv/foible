package server

import (
	"fmt"
	"log"

	"github.com/cfx-cv/herald/pkg/common"
	nats "github.com/nats-io/go-nats"
)

func (s *Server) subscribe() {
	s.conn.Subscribe(string(common.DijkstraErrors), func(msg *nats.Msg) {
		s.logs = append(s.logs, fmt.Sprintf("Topic=%s Message=%s", common.DijkstraErrors, string(msg.Data)))
	})
	s.conn.Subscribe(string(common.NamiErrors), func(msg *nats.Msg) {
		s.logs = append(s.logs, fmt.Sprintf("Topic=%s Message=%s", common.DijkstraErrors, string(msg.Data)))
	})
	s.conn.Subscribe(string(common.SuezErrors), func(msg *nats.Msg) {
		s.logs = append(s.logs, fmt.Sprintf("Topic=%s Message=%s", common.DijkstraErrors, string(msg.Data)))
	})

	if err := s.conn.Flush(); err != nil {
		log.Fatal(err)
	}
}
