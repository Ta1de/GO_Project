package main

import (
	"log"
	"network/Server"
)

func main() {
	port := Server.PortFlags()
	router := Server.NewRouter()
	server := Server.NewServer(port, router)
	if err := server.Run(port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
