package repositories

import (
	"hj_service/datasource"
	"hj_service/models"
)

var db = datasource.GetDb()

type UserRepository interface {
	GetUserList()
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (u userRepository) GetUserList() {
	/*var sql="select * from t_user"
	results, _ := datasource.GetDb().QueryInterface(sql)
	if b,err:=json.Marshal(results);err==nil{
		str:=string(b[:])
		println(str)
	}*/
	//var user []models.User;
	//datasource.GetDb().ID(1).Get(&user)
	//println(user.Phone)

	user := new(models.User)
	user.Phone = "1234"
	user.Id = 2
	user.Password = "123"
	user.Name = "lisi"
	datasource.GetDb().Insert(&models.User{Id: 2, Name: "zz", Password: "123", Phone: "13"})
}
