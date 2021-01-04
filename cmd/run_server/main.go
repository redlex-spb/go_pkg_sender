package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go_pkg_sender/internal/apps/server"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	gin.SetMode(gin.DebugMode)

	listenAddr := "8080"

	// create a router and server
	router := server.NewRouter()
	s := server.NewServer(
		listenAddr,
		(server.Middlewares{
			server.Tracing(func() string { return fmt.Sprintf("%d", time.Now().UnixNano()) }),
		}).Apply(router),
	)

	// run our server
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

}
