package contollers

import (
	"net/http"

	"github.com/Vedant-Jathar/Gin_Project/internal/models"
	"github.com/Vedant-Jathar/Gin_Project/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func (a *AuthController) InitController(authSrv services.AuthService) *AuthController {
	return &AuthController{
		authService: authSrv,
	}
}

func (a *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/login")
	routes.POST("/register")
}

func (a *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody models.AuthUser

		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err.Error(),
			})
		}
	}
}
