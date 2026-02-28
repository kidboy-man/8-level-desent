package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EchoController struct{}

func NewEchoController() *EchoController {
	return &EchoController{}
}

func (ctrl *EchoController) Echo(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid JSON body",
		})
		return
	}
	c.JSON(http.StatusOK, body)
}
