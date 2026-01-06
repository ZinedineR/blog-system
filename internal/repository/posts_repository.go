package repository

import (
	"blog-system/internal/entity"
)

type PostSQLRepo struct {
	Repository[entity.Posts]
}

func NewPostSQLRepository() PostRepository {
	return &PostSQLRepo{}
}
