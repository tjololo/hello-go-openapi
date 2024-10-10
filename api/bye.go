package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get Bye
// @Description get a bye message
// @Accept  plain/txt
// @Produce  plain/txt
// @Param   name     query    string     false        "Some Name"
// @Success 200 {string} string	"ok"
// @Router /bye [get]
func GetBye(c *gin.Context) {
	if c.Request.URL.Query().Has("name") {
		c.String(http.StatusOK, fmt.Sprintf("Bye %s", c.Request.URL.Query().Get("name")))
		return
	}
	c.String(http.StatusOK, "Bye World")
}
