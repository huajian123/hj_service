package services

import "hj_service/models"

type UserService interface {
	GetUserList(m map[string]interface{}) (result models.Result)
}
