package migration

import (
	"todoproject-be/src/database"
	"todoproject-be/src/models"
)

func DoMigration() {
	if (!database.Db.Migrator().HasTable(&models.User{})) {
		database.Db.AutoMigrate(&models.User{})
	}
	if (!database.Db.Migrator().HasTable(&models.Tip{})) {
		database.Db.AutoMigrate(&models.Tip{})
	}
}
