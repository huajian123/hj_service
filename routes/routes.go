package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	home_cotrl "hj_service/controller/home"
)

func App(app *iris.Application) {

	mvc.Configure(app.Party("/home"), func(context *mvc.Application) {
		context.Handle(home_cotrl.NewProjectCategoryController())
	})
}
