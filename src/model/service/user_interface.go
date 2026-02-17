package service

import (
	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/KevenAbraham/first-crud-go/src/model"
)

type userDomainService struct {}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *error_mapping.RestErr
	UpdateUser(string, model.UserDomainInterface) *error_mapping.RestErr
	FindUser(string) (*model.UserDomainInterface, *error_mapping.RestErr)
	DeleteUser(string) *error_mapping.RestErr
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
