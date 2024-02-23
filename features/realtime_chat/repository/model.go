package repository

import (
	"institute/features/realtime_chat"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) realtime_chat.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []realtime_chat.Realtime_chat {
	var realtime_chats []realtime_chat.Realtime_chat

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&realtime_chats)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return realtime_chats
}

func (mdl *model) Insert(newRealtime_chat realtime_chat.Realtime_chat) int64 {
	result := mdl.db.Create(&newRealtime_chat)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newRealtime_chat.ID)
}

func (mdl *model) SelectByID(realtime_chatID int) *realtime_chat.Realtime_chat {
	var realtime_chat realtime_chat.Realtime_chat
	result := mdl.db.First(&realtime_chat, realtime_chatID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &realtime_chat
}

func (mdl *model) Update(realtime_chat realtime_chat.Realtime_chat) int64 {
	result := mdl.db.Save(&realtime_chat)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(realtime_chatID int) int64 {
	result := mdl.db.Delete(&realtime_chat.Realtime_chat{}, realtime_chatID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}