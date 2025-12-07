package contollers

import (
	"github.com/Vedant-Jathar/Gin_Project/internal/models"
	"github.com/Vedant-Jathar/Gin_Project/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func (u *UserController) InitUserControllerRoutes(router *gin.Engine) {
	routes := router.Group("/user")
	routes.GET("/", u.GetUsers())
	routes.POST("/", u.CreateUser())
}

func (u *UserController) NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		users, err := u.UserService.GetUsers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"error":  err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"users":  users,
		})
	}
}

func (u *UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody models.User

		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err.Error(),
			})
			return
		}

		userId := u.UserService.CreateUser(&reqBody)

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User created",
			"userId":  userId,
		})
	}
}
