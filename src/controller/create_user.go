package controller

import (
	"fmt"
	"net/http"

	"github.com/MatheusJusto/go-crud/src/config/rest_err"
	"github.com/MatheusJusto/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorect field, error%s\any", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation successfull"),
	})
}
