package model

import (
	"fmt"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser() *error_mapping.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	ud.EncryptPassword()

	fmt.Println()
	return nil
}
