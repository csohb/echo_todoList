package migration

import (
	"Todo_List/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() *gorm.DB {
	// check your db info
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.TaskModel{}); err != nil {
		panic(err)
	}

	return db
}
