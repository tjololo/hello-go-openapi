package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get Hello
// @Description get a hello message
// @Accept  text/plain
// @Produce  text/plain
// @Param   name     query    string     false        "Some Name"
// @Success 200 {string} string	"ok"
// @Router /hello [get]
func GetHello(c *gin.Context) {
	if c.Request.URL.Query().Has("name") {
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", c.Request.URL.Query().Get("name")))
		return
	}
	c.String(http.StatusOK, "Hello World")
}
