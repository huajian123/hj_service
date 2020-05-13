package home_dao

import (
	"github.com/xormplus/xorm"
	"hj_service/models"
	home_model "hj_service/models/home"
	"log"
)

type ProjectCategoryDao struct {
	engine *xorm.Engine
}

func NewProjectCategoryDao(engine *xorm.Engine) *ProjectCategoryDao {
	return &ProjectCategoryDao{
		engine: engine,
	}
}

func (p *ProjectCategoryDao) GetProjectCategoryList(param models.SearchParam) []home_model.ProjectCategory {
	dataList := []home_model.ProjectCategory{}
	err := p.engine.Find(&dataList)
	if err != nil {
		log.Fatalln(err)
	}
	return dataList
}
