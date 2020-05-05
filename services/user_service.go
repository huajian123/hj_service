package services

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/xormplus/xorm"
	"hj_service/dao"
	"hj_service/datasource"
)

type UserService interface {
	GetUserList() int
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.GetDb()),
	}
}

func (u userService) GetUserList() int {
	return u.dao.GetUserList()
}
