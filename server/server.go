package server

import (
	log "github.com/cihub/seelog"
	L "github.com/vrecan/life"
)

// Server server struct
type Server struct {
	*L.Life
}

// NewServer creates a new server
func NewServer(c *Config) (s *Server) {
	log.Info("Config: ", c)
	s = &Server{
		Life: L.NewLife(),
	}
	s.SetRun(s.run)
	return s
}

func (s Server) run() {
	log.Info("Server Running")

	for {
		select {
		case <-s.Life.Done:
			log.Info("Server Successfully shut down")
			s.Life.WGDone()
			return
		}
	}
}

// Close satisfies io.Closer interface
func (s Server) Close() error {
	s.Life.Close()
	return nil
}
