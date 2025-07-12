package initializers

import (
	"github.com/thampaponn/learn-go/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
