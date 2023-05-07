package user

import "go_web_learning/model"

type Service interface {
	Login(id, password string) (*model.User, error)
	Register(in *model.User) (*model.User, error)
}
