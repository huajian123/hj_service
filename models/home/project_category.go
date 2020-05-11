package home_model

type ProjectCategory struct {
	Id            int64
	CategoryTitle string `xorm:"varchar(24)"`
}
