package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidboy-man/8-level-desent/app/models"
	"github.com/kidboy-man/8-level-desent/app/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) GenerateToken(c *gin.Context) {
	var req models.TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	token, err := ctrl.authService.GenerateToken(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{Token: token})
}
