package models

type User struct {
	Id       int64
	Name     string `xorm:"varchar(24)"`
	Phone    string `xorm:"varchar(24)"`
	Password string `xorm:"varchar(24)"`
}
