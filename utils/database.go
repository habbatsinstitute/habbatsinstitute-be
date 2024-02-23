package utils

import (
	"context"
	"fmt"
	"institute/config"
	"institute/features/auth"
	"institute/features/course"
	"institute/features/item"
	"institute/features/news"
	"institute/features/realtime_chat"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := config.LoadDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(auth.User{}, course.Course{}, news.News{}, news.Category{}, item.Item{}, realtime_chat.Realtime_chat{})
}

func ConnectMongo() *mongo.Database {
	config := config.LoadMongoConfig()

	clientOptions := options.Client()
	clientOptions.ApplyURI(config.MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil
	}
	return client.Database(config.MONGO_DB_NAME)
}