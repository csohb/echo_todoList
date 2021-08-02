package migration

import (
	"Todo_List/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() *gorm.DB {
	dsn := "test2:Cjswo.123@tcp(im.plea.kr:13306)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.TaskModel{}); err != nil {
		panic(err)
	}

	return db
}
