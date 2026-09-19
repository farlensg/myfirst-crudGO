package view

import (
	"github.com/farlensg/myfirst-crudGO/src/controller/model/response"
	"github.com/farlensg/myfirst-crudGO/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
