package services

import "github.com/KiraiChang/fingo/domain/entities"

type IUserService interface {
	Add(user *entities.NewUser)
	Get(filter *entities.UserFilter) *[]*entities.User
}
