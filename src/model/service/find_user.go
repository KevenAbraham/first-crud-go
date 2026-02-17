package service

import (
	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/KevenAbraham/first-crud-go/src/model"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *error_mapping.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	return nil, nil
}
