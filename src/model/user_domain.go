package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
)

type (
	UserDomain struct {
		Email    string
		Password string
		Name     string
		Age      int8
	}

	UserDomainInterface interface {
		CreateUser() *error_mapping.RestErr
		UpdateUser(string) *error_mapping.RestErr
		FindUser(string) (*UserDomain, *error_mapping.RestErr)
		DeleteUser(string) (*UserDomain, *error_mapping.RestErr)
	}
)

func NewUserDomain(
	email, password, name string,
	age int8,
) *UserDomain {
	return &UserDomain{
		email, password, name, age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
