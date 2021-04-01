package services

import (
	"github.com/KiraiChang/fingo/domain/entities"
	"github.com/KiraiChang/fingo/domain/interfaces/repos"
)

type UserService struct {
	repo *repos.IUserRepo
}

func (service UserService) Add(user *entities.NewUser) {
	(*service.repo).Add(user)
}

func (service UserService) Get(filter *entities.UserFilter) {
	(*service.repo).Get(filter)
}
