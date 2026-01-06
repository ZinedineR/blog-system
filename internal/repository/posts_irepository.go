package repository

import (
	"blog-system/internal/entity"
)

type PostRepository interface {
	// Example operations
	CommonQuery[entity.Posts]
}
