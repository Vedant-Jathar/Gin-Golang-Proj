package contollers

import (
	"fmt"
	"net/http"
	"strconv"
	// middleware "github.com/Vedant-Jathar/Gin_Project/internal/midlleware"
	"github.com/Vedant-Jathar/Gin_Project/internal/models"
	"github.com/Vedant-Jathar/Gin_Project/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (u *UserController) InitUserControllerRoutes(router *gin.Engine) {
	routes := router.Group("/user")

	routes.GET("/", u.GetUsers())
	routes.GET("/:id", u.GetUserById())
	routes.POST("/", u.CreateUser())
	routes.PUT("/:id", u.UpdateUser())
	routes.DELETE("/:id", u.DeleteUser())
}

func (u *UserController) NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println(c.Get("userId"))

		users, err := u.UserService.GetUsers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"users":  users,
		})
	}
}

func (u *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numId, err1 := strconv.Atoi(id)

		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err1.Error(),
			})
			return
		}

		user, err := u.UserService.GetUserById(numId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"user":    user,
			"message": "User fetched",
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

		userId, err3 := u.UserService.CreateUser(&reqBody)

		if err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err3,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User created",
			"userId":  userId,
		})
	}
}

func (u *UserController) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody models.User

		id := c.Param("id")

		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err.Error(),
			})
			return
		}

		fmt.Println("-------reqBody", reqBody)

		numId, err1 := strconv.Atoi(id)

		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err1.Error(),
			})
			return
		}

		data := reqBody

		err2 := u.UserService.UpdateUser(reqBody, data, numId)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err2.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User updated",
		})

	}
}

func (u *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numId, err1 := strconv.Atoi(id)

		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err1.Error(),
			})
		}

		if err := u.UserService.DeleteUser(numId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"error":  err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "User deleted",
		})
	}
}
