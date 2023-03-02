package service

import (
	"admin/constant"
	"admin/models"
	"admin/protocol"
	"admin/service/dto"
	"errors"

	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/logger"
)

type ImMessageService struct {
	service.Service
}

const NULL_ID int = 0

func (e *ImMessageService) GetMessages(message *dto.ImMessageRequest, list *[]dto.ImMessageResponse) error {
	//单聊
	if message.MessageType == constant.MESSAGE_TYPE_USER {

		var friend *models.SysUser
		e.Orm.First(&friend, "user_id = ?", message.Id)
		if NULL_ID == friend.UserId {
			return errors.New("用户不存在")
		}

		e.Orm.Raw("SELECT m.id, m.from_user_id, m.to_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username  FROM im_message AS m LEFT JOIN sys_user AS u ON m.from_user_id = u.user_id LEFT JOIN sys_user AS to_user ON m.to_id = to_user.user_id WHERE from_user_id IN (?, ?) AND to_id IN (?, ?) order by m.created_at",
			e.Identity.UserId, friend.UserId, e.Identity.UserId, friend.UserId).Scan(&list)

		return nil
	}
	//群聊
	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		return e.fetchGroupMessage(message.Id, list)
	}

	return errors.New("不支持查询类型")
}

func (e *ImMessageService) fetchGroupMessage(groupId int, list *[]dto.ImMessageResponse) error {
	var group models.ImGroup
	e.Orm.First(&group, "id = ?", groupId)
	if group.Id <= 0 {
		return errors.New("群组不存在")
	}

	e.Orm.Raw("SELECT m.id, m.from_user_id, m.to_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar FROM im_message AS m LEFT JOIN sys_user AS u ON m.from_user_id = u.user_id WHERE m.message_type = 2 AND m.to_id = ?",
		group.Id).Scan(&list)

	return nil
}

func (e *ImMessageService) SaveMessage(message protocol.Message) {
	var fromUser models.SysUser
	e.Orm.Find(&fromUser, "user_id = ?", message.From)
	if NULL_ID == fromUser.UserId {
		logger.Errorf("SaveMessage not find from user", fromUser.UserId)
		return
	}

	var toId int = 0
	if message.MessageType == constant.MESSAGE_TYPE_USER {
		var toUser models.SysUser
		e.Orm.Find(&toUser, "user_id = ?", message.To)
		if NULL_ID == toUser.UserId {
			return
		}
		toId = toUser.UserId
	}

	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		var group models.ImGroup
		e.Orm.Find(&group, "id = ?", message.To)
		if NULL_ID == group.Id {
			return
		}
		toId = group.Id
	}

	saveMessage := models.ImMessage{
		FromUserId:  fromUser.UserId,
		ToId:        toId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	e.Orm.Save(&saveMessage)
}
