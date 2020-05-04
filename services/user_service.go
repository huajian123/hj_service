package services

import (
	"hj_service/models"
)

type UserService interface {
	GetUserList() (result models.Result)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u userService) GetUserList() (result models.Result) {
	result.Data = "User"
	result.Code = 0
	result.Msg = "SUCCESS"
	return result
}
