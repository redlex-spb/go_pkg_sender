package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", index)
	r.StaticFS("/pkg/", http.Dir("./ps4_pkg"))

	return r
}
