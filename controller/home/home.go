package home_cotrl

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"hj_service/models"
	home_service "hj_service/services/home"
)

type ProjectCategoryController struct {
	Ctx     iris.Context
	Service home_service.ProjectCategoryService
}

func NewProjectCategoryController() *ProjectCategoryController {
	return &ProjectCategoryController{Service: home_service.NewProjectCategoryService()}
}

func (projectCategoryController *ProjectCategoryController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/queryCategory", "QueryCategory")
}

func (pc *ProjectCategoryController) QueryCategory() mvc.Result {
	data := pc.Service.GetProjectCategoryList()
	return mvc.Response{
		Object: models.NewResult(data, 0),
	}
}
