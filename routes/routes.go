package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"hj_service/controller"
)

func App(app *iris.Application) {
	mvc.Configure(app.Party("/user"), func(context *mvc.Application) {
		context.Handle(new(controller.UserController))
	})
}
