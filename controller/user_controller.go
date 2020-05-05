package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"hj_service/models"
	"hj_service/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

func NewUserCotroller() *UserController {
	return &UserController{Service: services.NewUserService()}
}

func (userController *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/query", "UserInfo")
	a.Handle("GET", "/app", "Aaa")
}

func (uc *UserController) UserInfo() (result models.Result) {
	return uc.Service.GetUserList()
}
func (uc *UserController) Aaa() int64 {
	return 123
}
