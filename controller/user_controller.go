package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"hj_service/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

func (userController *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/query", "UserInfo")
}

func (uc *UserController) UserInfo() mvc.Result {
	//todo
	iris.New().Logger().Info(" user info query ")
	return mvc.Response{}
}
