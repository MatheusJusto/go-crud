package controller

import (
	"net/http"

	"github.com/MatheusJusto/go-crud/src/config/logger"
	"github.com/MatheusJusto/go-crud/src/config/validation"
	"github.com/MatheusJusto/go-crud/src/controller/model/request"
	"github.com/MatheusJusto/go-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createdUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "createdUser"),
	)

	c.JSON(http.StatusCreated, response)
}
