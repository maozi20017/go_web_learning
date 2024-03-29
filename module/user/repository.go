package user

import "go_web_learning/model"

type Repository interface {
	GetUserList(map[string]interface{}) ([]*model.User, error)
	GetUser(in *model.User) (*model.User, error)
	CreateUser(in *model.User) (*model.User, error)
	ModifyUser(in *model.User, data map[string]interface{}) (*model.User, error)
	DeleteUser(in *model.User) error
}
