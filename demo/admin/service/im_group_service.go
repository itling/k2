package service

import (
	"admin/models"
	"admin/service/dto"
	"errors"

	"github.com/kingwel-xie/k2/common/service"
)

type ImGroupService struct {
	service.Service
}

func (e *ImGroupService) GetGroups(userId string) ([]dto.GroupResponse, error) {

	var queryUser *models.SysUser
	e.Orm.First(&queryUser, "user_id = ?", userId)

	if queryUser.UserId <= 0 {
		return nil, errors.New("用户不存在")
	}

	var groups []dto.GroupResponse

	e.Orm.Raw("SELECT g.id AS group_id, g.created_at, g.name, g.notice FROM im_group_members AS gm LEFT JOIN im_groups AS g ON gm.group_id = g.id WHERE gm.user_id = ?",
		queryUser.UserId).Scan(&groups)

	return groups, nil
}

func (e *ImGroupService) SaveGroup(userId string, group models.ImGroup) {
	var fromUser models.SysUser
	e.Orm.Find(&fromUser, "user_id = ?", userId)
	if fromUser.UserId <= 0 {
		return
	}

	group.UserId = fromUser.UserId
	e.Orm.Save(&group)

	groupMember := models.ImGroupMember{
		UserId:   fromUser.UserId,
		GroupId:  group.Id,
		Nickname: fromUser.NickName,
		Mute:     0,
	}
	e.Orm.Save(&groupMember)
}

func (e *ImGroupService) GetUserIdByGroupId(groupId int) []models.SysUser {
	var group models.ImGroup
	e.Orm.First(&group, "id = ?", groupId)
	if group.Id <= 0 {
		return nil
	}

	var users []models.SysUser
	e.Orm.Raw("SELECT u.user_id, u.avatar, u.username FROM im_groups AS g JOIN im_group_members AS gm ON gm.group_id = g.id JOIN users AS u ON u.user_id = gm.user_id WHERE g.id = ?",
		group.Id).Scan(&users)
	return users
}

func (e *ImGroupService) JoinGroup(groupId, userId string) error {
	var user models.SysUser
	e.Orm.First(&user, "user_id = ?", userId)
	if user.UserId <= 0 {
		return errors.New("用户不存在")
	}

	var group models.ImGroup
	e.Orm.First(&group, "id = ?", groupId)
	if user.UserId <= 0 {
		return errors.New("群组不存在")
	}

	var groupMember models.ImGroupMember
	e.Orm.First(&groupMember, "user_id = ? and group_id = ?", user.UserId, group.Id)
	if groupMember.Id > 0 {
		return errors.New("已经加入该群组")
	}
	nickname := user.NickName
	if nickname == "" {
		nickname = user.Username
	}
	groupMemberInsert := models.ImGroupMember{
		UserId:   user.UserId,
		GroupId:  group.Id,
		Nickname: nickname,
		Mute:     0,
	}
	e.Orm.Save(&groupMemberInsert)

	return nil
}
