package home

type ProjectCategory struct {
	Id            int64
	CategoryTitle string `xorm:"varchar(24)"`
}
