package initializers

import "jwt/models"

func MigrateTables() {
	DB.AutoMigrate(&models.User{})
}
