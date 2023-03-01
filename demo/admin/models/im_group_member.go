package models

import (
	"time"
	"github.com/kingwel-xie/k2/common/models"
)

type ImGroupMember struct {
	Id       int    `json:"id" gorm:"primarykey"`
	GroupId  int    `json:"groupId" gorm:"index;comment:'群组ID'"`
	UserId   int    `json:"userId" gorm:"index;comment:'用户ID'"`
	Nickname string `json:"nickname" gorm:"type:varchar(350);comment:'昵称"`
	Mute     int16  `json:"mute" gorm:"comment:'是否禁言'"`
	models.ModelTime
}

func (ImGroupMember) TableName() string {
	return "im_group_member"
}


type ImGroupRequest struct {
	UserId        int      `gorm:"column:user_id;primary_key;size:64;comment:用户id"`
	GroupId       int       `gorm:"column:group_id;primary_key;size:64;comment:群组id"`
	HandleResult  int32     `gorm:"column:handle_result;comment:处理结果:1等待验证 2通过 3拒绝"`
	ReqMsg        string    `gorm:"column:req_msg;size:1024;comment:申请加入群组消息"`
	HandledMsg    string    `gorm:"column:handle_msg;size:1024;comment:处理附带消息"`
	ReqTime       time.Time `gorm:"column:req_time;comment:申请加入群组时间"`
	HandleUserId  int       `gorm:"column:handle_user_id;size:64;comment:处理人id"`
	HandledTime   time.Time `gorm:"column:handle_time;comment:处理时间"`
	InviterUserId int       `gorm:"column:inviter_user_id;size:64;comment:邀请人id"`
	models.ModelTime
}

func (ImGroupRequest) TableName() string {
	return "im_group_request"
}
