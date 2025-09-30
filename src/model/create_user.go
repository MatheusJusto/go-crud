package model

import (
	"fmt"

	"github.com/MatheusJusto/go-crud/src/config/logger"
	"github.com/MatheusJusto/go-crud/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init create model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println(ud)

	return nil
}
