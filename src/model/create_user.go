package model

import (
	"fmt"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *error_mapping.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println()
	return nil
}
