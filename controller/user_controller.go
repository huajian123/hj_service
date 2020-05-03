package controller

import (
	"github.com/kataras/iris/v12"
	"hj_service/services"
)

type LoginController struct {
	Ctx     iris.Context
	Service services.UserService
}
