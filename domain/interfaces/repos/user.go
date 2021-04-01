package repos

import "github.com/KiraiChang/fingo/domain/entities"

type IUserRepo interface {
	Add(user *entities.NewUser)
	Get(filter *entities.UserFilter) *[]*entities.User
	Page(page *entities.UserPagination) *[]*entities.User
}
