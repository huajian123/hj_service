package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"hj_service/controller"
	home_cotrl "hj_service/controller/home"
)

func App(app *iris.Application) {
	mvc.Configure(app.Party("/user"), func(context *mvc.Application) {
		context.Handle(controller.NewUserCotroller())
	})

	mvc.Configure(app.Party("/home"), func(context *mvc.Application) {
		context.Handle(home_cotrl.NewProjectCategoryController())
	})
}
