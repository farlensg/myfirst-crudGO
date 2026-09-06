package model

import (
	"github.com/farlensg/myfirst-crudGO/src/configuration/logger"
	"github.com/farlensg/myfirst-crudGO/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	return nil
}
