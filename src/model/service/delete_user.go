package service

import (
	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(string) *error_mapping.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
