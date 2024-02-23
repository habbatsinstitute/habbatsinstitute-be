package usecase

import (
	"institute/features/realtime_chat"
	"institute/features/realtime_chat/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model realtime_chat.Repository
}

func New(model realtime_chat.Repository) realtime_chat.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResRealtime_chat {
	var realtime_chats []dtos.ResRealtime_chat

	realtime_chatsEnt := svc.model.Paginate(page, size)

	for _, realtime_chat := range realtime_chatsEnt {
		var data dtos.ResRealtime_chat

		if err := smapping.FillStruct(&data, smapping.MapFields(realtime_chat)); err != nil {
			log.Error(err.Error())
		} 
		
		realtime_chats = append(realtime_chats, data)
	}

	return realtime_chats
}

func (svc *service) FindByID(realtime_chatID int) *dtos.ResRealtime_chat {
	res := dtos.ResRealtime_chat{}
	realtime_chat := svc.model.SelectByID(realtime_chatID)

	if realtime_chat == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(realtime_chat))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newRealtime_chat dtos.InputRealtime_chat) *dtos.ResRealtime_chat {
	realtime_chat := realtime_chat.Realtime_chat{}
	
	err := smapping.FillStruct(&realtime_chat, smapping.MapFields(newRealtime_chat))
	if err != nil {
		log.Error(err)
		return nil
	}

	realtime_chatID := svc.model.Insert(realtime_chat)

	if realtime_chatID == -1 {
		return nil
	}

	resRealtime_chat := dtos.ResRealtime_chat{}
	errRes := smapping.FillStruct(&resRealtime_chat, smapping.MapFields(newRealtime_chat))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resRealtime_chat
}

func (svc *service) Modify(realtime_chatData dtos.InputRealtime_chat, realtime_chatID int) bool {
	newRealtime_chat := realtime_chat.Realtime_chat{}

	err := smapping.FillStruct(&newRealtime_chat, smapping.MapFields(realtime_chatData))
	if err != nil {
		log.Error(err)
		return false
	}

	newRealtime_chat.ID = realtime_chatID
	rowsAffected := svc.model.Update(newRealtime_chat)

	if rowsAffected <= 0 {
		log.Error("There is No Realtime_chat Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(realtime_chatID int) bool {
	rowsAffected := svc.model.DeleteByID(realtime_chatID)

	if rowsAffected <= 0 {
		log.Error("There is No Realtime_chat Deleted!")
		return false
	}

	return true
}