package routes

import (
	"github.com/KevenAbraham/first-crud-go/src/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", handler.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", handler.FindUserByEmail)
	r.POST("/createUser", handler.CreateUser)
	r.PUT("/updateUser/:userId", handler.UpdateUser)
	r.DELETE("/deleteUser/:userId", handler.DeleteUser)
}
