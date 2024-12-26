package Server

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	Port       int
	httpServer http.Server
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		Port: port,
		httpServer: http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: handler,
		},
	}
}

func (s *Server) Run(port int) error {
	log.Printf("server is running on port: %d", port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

func PortFlags() int {
	portFlag := flag.Int("P", 8080, "Port for the server")
	flag.Parse()

	var port int = 8080 // Значение по умолчанию
	args := flag.Args()

	if len(args) > 0 {
		portInput := args[0]
		var err error
		port, err = strconv.Atoi(portInput)
		if err != nil || port <= 0 {
			log.Fatalf("Неверный порт: %v", portInput)
		}
	} else if *portFlag != 8080 {
		port = *portFlag
	}
	return port
}
