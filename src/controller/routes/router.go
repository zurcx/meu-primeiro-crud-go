package routes

import (
	"github.com/gin-gonic/gin"
	"ithub.com/zurcx/huncoding/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser, controller.DeleteUser)
	r.DELETE("/deleteUser/:userId")
}
