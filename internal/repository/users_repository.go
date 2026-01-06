package repository

import (
	"blog-system/internal/entity"
)

type UserSQLRepo struct {
	Repository[entity.Users]
}

func NewUserSQLRepository() UserRepository {
	return &UserSQLRepo{}
}
