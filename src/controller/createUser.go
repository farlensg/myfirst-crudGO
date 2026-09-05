package controller

import (
	"net/http"

	"github.com/farlensg/myfirst-crudGO/src/configuration/logger"
	"github.com/farlensg/myfirst-crudGO/src/configuration/validation"
	"github.com/farlensg/myfirst-crudGO/src/controller/model/request"
	"github.com/farlensg/myfirst-crudGO/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("jorney", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("jorney", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User create successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
