package dao

import (
	"github.com/xormplus/xorm"
	"hj_service/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (u *UserDao) GetUserList() int {
	u.engine.InsertOne(&models.User{Password: "55", Id: 6, Phone: "5"})
	return 666
}
