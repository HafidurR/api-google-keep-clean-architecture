package config

import (
	"fmt"

	"api-google-keep/core/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&entity.User{}, &entity.Note{})
 
	return db
}