package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
