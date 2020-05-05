package dao

import "github.com/xormplus/xorm"

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (u *UserDao) GetUserList() int {
	return 666
}
