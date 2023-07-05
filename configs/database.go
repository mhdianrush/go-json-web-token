package configs

import (
	"github.com/mhdianrush/go-json-web-token/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	logger := logrus.New()

	db, err := gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/go_jwt?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	DB = db

	logger.Println("Database Connected")
}
