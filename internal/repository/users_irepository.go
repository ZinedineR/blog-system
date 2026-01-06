package repository

import (
	"blog-system/internal/entity"
)

type UserRepository interface {
	// Example operations
	CommonQuery[entity.Users]
}
