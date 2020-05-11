package home

import "github.com/xormplus/xorm"

type ProjectCategoryDao struct {
	engine *xorm.Engine
}

func NewProjectCategoryDao(engine *xorm.Engine) *ProjectCategoryDao {
	return &ProjectCategoryDao{
		engine: engine,
	}
}
