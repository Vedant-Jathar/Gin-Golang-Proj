// package main

// import (
// 	contollers "github.com/Vedant-Jathar/Gin_Project/internal/controllers"
// 	"github.com/Vedant-Jathar/Gin_Project/internal/database"
// 	"github.com/Vedant-Jathar/Gin_Project/internal/models"
// 	"github.com/Vedant-Jathar/Gin_Project/internal/services"
// 	"github.com/gin-gonic/gin"
// 	// student "github.com/Vedant-Jathar/Gin_Project/internal/types"
// )

// func main() {
// 	router := gin.Default()

// 	db := database.ConnectDb()

// 	if err := db.AutoMigrate(&models.User{},&models.AuthUser{}); err != nil {
// 		panic(err)
// 	}

// 	UserSrv := services.UserService{}
// 	newUserSrv := UserSrv.NewUserService(db)

// 	UserCtlr := contollers.UserController{}
// 	newUserCtlr := UserCtlr.NewUserController(*newUserSrv)
// 	newUserCtlr.InitUserControllerRoutes(router)

// 	// Auth controllers:

// 	authSrv := services.AuthService{}
// 	newAuthSrv := authSrv.InitAuthservice(db)

// 	authCtlr := contollers.AuthController{}
// 	newAuthCtlr := authCtlr.InitController(*newAuthSrv)
// 	newAuthCtlr.InitRoutes(router)

// 	router.Run(":8000")

// 	// router.GET("/ping", func(c *gin.Context) {
// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"status": true,
// 	// 		"method": c.Request.Method,
// 	// 	})
// 	// })

// 	// router.GET("/profile/:id", func(c *gin.Context) {
// 	// 	id := c.Param("id")
// 	// 	num_id, err := strconv.Atoi(id)

// 	// 	if err != nil {
// 	// 		c.JSON(500, gin.H{
// 	// 			"error": err.Error(),
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"id": num_id,
// 	// 	})
// 	// })

// 	// router.POST("/user", func(c *gin.Context) {

// 	// 	var reqBody student.Student

// 	// 	if err := c.BindJSON(&reqBody); err != nil {
// 	// 		c.JSON(http.StatusBadRequest, gin.H{
// 	// 			"error": err.Error(),
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	c.JSON(http.StatusCreated, gin.H{
// 	// 		"email": reqBody.Email,
// 	// 		"name":  reqBody.Name,
// 	// 	})

// 	// })

// 	// router.PUT("/user/:id", func(c *gin.Context) {
// 	// 	id := c.Param("id")

// 	// 	var reqBody student.Student

// 	// 	if err := c.BindJSON(&reqBody); err != nil {
// 	// 		c.JSON(http.StatusBadRequest, gin.H{
// 	// 			"status": false,
// 	// 			"err":    err.Error(),
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"id":      id,
// 	// 		"status":  true,
// 	// 		"reqBody": reqBody,
// 	// 	})
// 	// })

// }
