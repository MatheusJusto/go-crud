package controller

import (
	"fmt"
	"net/http"

	"github.com/MatheusJusto/go-crud/src/config/validation"
	"github.com/MatheusJusto/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation successfull"),
	})
}
