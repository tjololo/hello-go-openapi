package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get adieu
// @Description get a adieu message
// @Accept  text/plain
// @Produce  text/plain
// @Param   name     query    string     false        "Some Name"
// @Success 200 {string} string	"ok"
// @Router /adieu [get]
func GetBye(c *gin.Context) {
	if c.Request.URL.Query().Has("name") {
		c.String(http.StatusOK, fmt.Sprintf("Bye %s", c.Request.URL.Query().Get("name")))
		return
	}
	c.String(http.StatusOK, "Bye World")
}
