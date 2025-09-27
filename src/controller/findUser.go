package controller

import (
	"github.com/MatheusJusto/go-crud/src/config/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
