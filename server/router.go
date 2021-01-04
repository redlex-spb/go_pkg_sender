package server

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	// routes
	router.Handle("/", index())
	//router.Handle("/healthz", healthz())

	fileServer := http.FileServer(http.Dir("./ps4_pkg"))
	router.Handle("/pkg/", http.StripPrefix("/pkg/", fileServer))

	return router
}
