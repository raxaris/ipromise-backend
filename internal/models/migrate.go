package models

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
		&RefreshToken{},
		&Promise{},
	)
	if err != nil {
		panic("❌ Ошибка миграции: " + err.Error())
	}
}
