package services

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/xormplus/xorm"
	"hj_service/dao"
	"hj_service/datasource"
)

// 接口
type UserService interface {
	GetUserList() int
}

// 类
type userService struct {
	dao *dao.UserDao
}

// 构造函数
func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.GetDb()),
	}
}

// 类方法
func (u userService) GetUserList() int {
	return u.dao.GetUserList()
}
