package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented,
		ResponseWrapper{
			Status:   "error",
			Code:     http.StatusNotImplemented,
			Messages: []string{"Not Implemented"},
		})
}
