package controller

import "github.com/kataras/iris/v12"

type LoginController struct {
	Ctx     iris.Context
	Service service.LoginService
}
