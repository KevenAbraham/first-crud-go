package handler

import (
	"fmt"
	"log"

	"github.com/KevenAbraham/first-crud-go/src/configurations/validation"
	"github.com/KevenAbraham/first-crud-go/src/handler/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error: %s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
