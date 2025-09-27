package routes

import (
	"github.com/MatheusJusto/go-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/get-user-by-id/:user-id", controller.FindUserById)
	r.GET("/get-user-by-email/:user-email", controller.FindUserByEmail)
	r.POST("/create-user", controller.CreateUser)
	r.PUT("/update-user/:user-id", controller.UpdateUser)
	r.DELETE("/delete-user/:user-id", controller.DeleteUser)
}
