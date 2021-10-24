package DB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"go-server-template/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//数据库连接
func InitDBGorm() *gorm.DB {
	config := projectConfig.AppConfig
	DatabaseConfig := config.DatabaseConfig
	host := DatabaseConfig.HOST
	port := DatabaseConfig.PORT
	database := DatabaseConfig.DATABASE
	username := DatabaseConfig.USER_NAME
	password := DatabaseConfig.PASSWORD
	charset := DatabaseConfig.CHARSET
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Error("failed to connect database,err:" + err.Error())
	}

	// 自动同步库
	AutoMigrateDBGorm(db)

	return db

}
