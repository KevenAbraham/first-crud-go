package handler

import (
	"net/http"

	"github.com/KevenAbraham/first-crud-go/src/configurations/logger"
	"github.com/KevenAbraham/first-crud-go/src/configurations/validation"
	"github.com/KevenAbraham/first-crud-go/src/handler/model/request"
	"github.com/KevenAbraham/first-crud-go/src/handler/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateUser(ctx *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully", zap.String("journey", "createUser"))

	ctx.JSON(http.StatusOK, response)
}
