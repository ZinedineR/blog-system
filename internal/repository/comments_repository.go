package repository

import (
	"blog-system/internal/entity"
)

type CommentSQLRepo struct {
	Repository[entity.Comments]
}

func NewCommentSQLRepository() CommentRepository {
	return &CommentSQLRepo{}
}
