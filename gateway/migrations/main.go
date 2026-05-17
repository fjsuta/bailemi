package main

import (
	"fmt"
	"log"

	"github.com/bailemi/gateway/configs"
	"github.com/bailemi/gateway/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg := configs.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	log.Println("开始数据库迁移...")

	err = db.AutoMigrate(
		&model.User{},
		&model.UserProfile{},
		&model.Follow{},
		&model.VIPRecord{},
		&model.Artist{},
		&model.Album{},
		&model.Song{},
		&model.Lyric{},
		&model.Fingerprint{},
		&model.Genre{},
		&model.Playlist{},
		&model.PlaylistSong{},
		&model.Favorite{},
		&model.Comment{},
		&model.Like{},
		&model.Message{},
		&model.Report{},
		&model.PlayHistory{},
		&model.SearchHistory{},
		&model.UserBehaviorLog{},
		&model.Upload{},
	)
	if err != nil {
		log.Fatal("自动迁移失败:", err)
	}

	log.Println("数据库迁移完成!")

	seedGenres(db)
	seedAdminUser(db)

	log.Println("数据初始化完成!")
}

func seedGenres(db *gorm.DB) {
	var count int64
	db.Model(&model.Genre{}).Count(&count)
	if count > 0 {
		log.Println("流派数据已存在，跳过初始化")
		return
	}

	genres := []model.Genre{
		{Name: "流行", NameEn: strPtr("Pop"), SortOrder: 1},
		{Name: "摇滚", NameEn: strPtr("Rock"), SortOrder: 2},
		{Name: "民谣", NameEn: strPtr("Folk"), SortOrder: 3},
		{Name: "电子", NameEn: strPtr("Electronic"), SortOrder: 4},
		{Name: "古典", NameEn: strPtr("Classical"), SortOrder: 5},
		{Name: "爵士", NameEn: strPtr("Jazz"), SortOrder: 6},
		{Name: "说唱", NameEn: strPtr("Hip-Hop"), SortOrder: 7},
		{Name: "R&B", NameEn: strPtr("R&B"), SortOrder: 8},
		{Name: "轻音乐", NameEn: strPtr("Easy Listening"), SortOrder: 9},
		{Name: "古典", NameEn: strPtr("Classical"), SortOrder: 10},
	}

	for i := range genres {
		db.Create(&genres[i])
	}

	log.Println("流派数据初始化完成")
}

func seedAdminUser(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Where("role = ?", 9).Count(&count)
	if count > 0 {
		log.Println("管理员用户已存在，跳过初始化")
		return
	}

	admin := &model.User{
		Username:     "admin",
		PasswordHash: "$2a$10$N9qo8uLOickgx2ZMRZoMye.Gqgz9N0B8V.S6l6O0lLx1b3P5H5q6y",
		Role:        9,
		Status:      1,
	}
	db.Create(admin)

	log.Println("管理员用户初始化完成")
}

func strPtr(s string) *string {
	return &s
}
