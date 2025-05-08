package models

import (
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
		&RefreshToken{},
		&Promise{},
		&Microtask{},
		&Post{},
		&Attachment{},
		&Follower{},
		&Like{},
		&Badge{},
		&UserBadge{},
	)
	if err != nil {
		panic("❌ Ошибка миграции: " + err.Error())
	}
}
