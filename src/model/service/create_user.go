package service

import (
	"fmt"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/KevenAbraham/first-crud-go/src/model"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *error_mapping.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println()
	return nil
}
