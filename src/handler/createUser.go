package handler

import (
	"fmt"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/KevenAbraham/first-crud-go/src/handler/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := error_mapping.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error: %s\n", err.Error()))
		
		ctx.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}
