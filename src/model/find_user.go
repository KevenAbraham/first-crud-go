package model

import (
	"fmt"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser() *error_mapping.RestErr {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	ud.EncryptPassword()

	fmt.Println()
	return nil
}
