package controller

import (
	"fmt"
	"log"

	"github.com/farlensg/myfirst-crudGO/src/configuration/validation"
	"github.com/farlensg/myfirst-crudGO/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to bind json, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
