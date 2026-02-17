package handler

import (
	"net/http"

	"github.com/KevenAbraham/first-crud-go/src/configurations/logger"
	"github.com/KevenAbraham/first-crud-go/src/configurations/validation"
	"github.com/KevenAbraham/first-crud-go/src/handler/model/request"
	"github.com/KevenAbraham/first-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age,
	)
	if err := domain.CreateUser(); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.String(http.StatusOK, "")
}
