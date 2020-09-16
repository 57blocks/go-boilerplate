package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resource struct{}

func New() *Resource {
	return &Resource{}
}

func (rs *Resource) Ping(c *gin.Context) {
	type pong struct {
		Value string `json:"value"`
	}
	c.JSON(http.StatusOK, &pong{"pong"})
}
