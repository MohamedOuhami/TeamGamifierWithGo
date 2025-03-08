package initializers

import "github.com/MohamedOuhami/TeamGamifierWithGo/models"

func SyncDatabase() {
	// Migrate the schemas into the database to create the tables
	DB.AutoMigrate(&models.User{})
}
