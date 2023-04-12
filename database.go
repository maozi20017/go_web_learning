package main

import (
	"fmt"
	"go_web_learning/config"
	"go_web_learning/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("無法讀取配置檔案: %s", err)
	}

	var config config.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("無法解析配置檔案: %s", err)
	}

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Network, config.Server, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("使用 gorm 連線 DB 發生錯誤，原因為 %s", err)
	}

	if err := db.AutoMigrate(new(model.User)); err != nil { // 自動建立 User 資料表
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}

	db.FirstOrCreate(&model.User{ID: 1, Username: "10811212", Password: "10811212", IdentityID: 1})
	db.FirstOrCreate(&model.User{ID: 2, Username: "10811220", Password: "10811220", IdentityID: 2})
	db.FirstOrCreate(&model.User{ID: 3, Username: "10811225", Password: "ghost8797", IdentityID: 3})

	if err := db.AutoMigrate(new(model.Identity)); err != nil { // 自動建立 Role 資料表
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}

	db.FirstOrCreate(&model.Identity{ID: 1, Description: "一般使用者"})
	db.FirstOrCreate(&model.Identity{ID: 2, Description: "版主"})
	db.FirstOrCreate(&model.Identity{ID: 3, Description: "管理員"})

	return db, nil
}
