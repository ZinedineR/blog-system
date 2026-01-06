package migration

import (
	"blog-system/internal/entity"
	"blog-system/pkg/database"
)

func AutoMigration(CpmDB *database.Database) {
	CpmDB.MigrateDB(
		&entity.Users{},
		&entity.Posts{},
		&entity.Comments{})
	//&entity.SMSLog{}
}
