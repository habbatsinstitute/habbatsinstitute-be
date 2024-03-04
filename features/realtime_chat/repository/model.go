package repository

import (
	"context"
	realtimechat "institute/features/realtime_chat"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
	collection *mongo.Collection
}

func New(db *gorm.DB,collection *mongo.Collection) realtimechat.Repository {
	return &model {
		db: db,
		collection: collection,
	}
}

func (mdl *model) SaveChat(questionNReply realtimechat.Chat, userID int, recipientID int) error {
	var user realtimechat.User

	if err := mdl.collection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&user); err != nil {
		mdl.collection.InsertOne(context.Background(), realtimechat.PrivateChatHistory{
			SenderID: userID,
			RecipientID: recipientID,
			Message: []realtimechat.Chat{
				questionNReply,
			},
		})
	}else {
		if _, err := mdl.collection.UpdateOne(context.Background(), bson.M{"user_id": userID}, bson.M{"$push": bson.M{"question_reply": questionNReply}}); err != nil {
			return err
		}
	}
	return nil
}

func (mdl *model) SelectByID(userID int) *realtimechat.User {
	var user realtimechat.User
	result := mdl.db.Table("users").First(&user, userID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &user
}

func (mdl *model) TimeStamp() time.Time{
	timestamp ,_ := time.LoadLocation("Asia/Jakarta")

	return time.Now().In(timestamp)
}