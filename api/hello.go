package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHello(c *gin.Context) {
	fmt.Printf("Hello %v\n", c.Request.URL.Query().Get("name"))
	if c.Request.URL.Query().Has("name") {
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", c.Request.URL.Query().Get("name")))
		return
	}
	c.String(http.StatusOK, "Hello World")
}
