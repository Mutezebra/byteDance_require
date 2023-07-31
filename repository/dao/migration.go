package dao

import (
	"AlittleRequire/repository/model"
	"log"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Topic{}, &model.Post{})
	if err != nil {
		return
	}
	log.Println("迁移成功")
}
