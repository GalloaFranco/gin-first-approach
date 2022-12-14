package repositories

import (
	"github.com/GalloaFranco/gin-first-approach/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type IVideoRepository interface {
	Save(video entities.Video) entities.Video
	Update(video entities.Video)
	Delete(video entities.Video)
	FindAll() []entities.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() IVideoRepository {
	dsn := "host=localhost user=root password=postgres dbname=docker_pg_db port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect with database")
	}
	db.Logger.LogMode(logger.Info)
	// Automatically migrate your schema.
	if err := db.AutoMigrate(&entities.Video{}); err != nil { // This would create a new table in the DB
		log.Println(err.Error())
		panic("Failed to migrate database")
	}
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	postgresDB, _ := db.connection.DB()
	if err := postgresDB.Close(); err != nil {
		panic("Cannot close the database connection")
	}
}

func (db *database) Save(video entities.Video) entities.Video {
	db.connection.Create(&video)
	return video
}

func (db *database) Update(video entities.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entities.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entities.Video {
	var videos []entities.Video
	db.connection.Find(&videos)
	return videos
}
