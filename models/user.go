package models

type User struct {
	Id       int64
	Phone    string `xorm:"varchar(24)" json:"phone"`
	Password string `xorm:"varchar(24)" json:"password"`
}
