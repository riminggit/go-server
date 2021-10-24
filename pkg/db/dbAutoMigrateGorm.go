package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"go-server-template/model/user"
	"gorm.io/gorm"
)

func AutoMigrateDBGorm(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})

	// db.Create(&userModel.User{Account: "123456789"})
}
