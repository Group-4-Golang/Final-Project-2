package helper

import (
	"hacktiv-final2/comment"
	"hacktiv-final2/photo"
	"hacktiv-final2/socialmedia"
	"hacktiv-final2/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	cfg := LoadConfig()
	dsn := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port + ")/" + cfg.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{}, &socialmedia.SocialMedia{}, &photo.Photo{}, &comment.Comment{})
	return db
}
