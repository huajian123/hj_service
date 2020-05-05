package services

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/xormplus/xorm"
	"hj_service/models"
	"hj_service/repositories"
)

type UserService interface {
	GetUserList() (result models.Result)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u userService) GetUserList() (result models.Result) {
	// 写sql获取结果返回
	repositories.NewUserRepository().GetUserList()

	result.Data = "User"
	result.Code = 0
	result.Msg = "SUCCESS"
	return result
}
