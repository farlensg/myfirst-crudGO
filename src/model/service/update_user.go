package service

import (
	"github.com/farlensg/myfirst-crudGO/src/configuration/rest_err"
	"github.com/farlensg/myfirst-crudGO/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
