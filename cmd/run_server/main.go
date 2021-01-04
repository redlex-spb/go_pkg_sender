package main

import (
	"fmt"
	"go_pkg_sender/server"
	"log"
	"os"
	"time"
)

func main() {
	listenAddr := "8080"

	// create a logger, router and run_server
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	router := server.NewRouter()
	s := server.NewServer(
		listenAddr,
		(server.Middlewares{
			server.Logging(logger), server.Tracing(func() string { return fmt.Sprintf("%d", time.Now().UnixNano()) }),
		}).Apply(router),
		logger,
	)

	// run our run_server
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
