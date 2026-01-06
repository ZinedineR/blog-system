package repository

import (
	"blog-system/internal/entity"
)

type CommentRepository interface {
	// Example operations
	CommonQuery[entity.Comments]
}
