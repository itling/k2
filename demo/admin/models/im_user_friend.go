package models

import (
	"time"
	"github.com/kingwel-xie/k2/common/models"
)

type ImUserFriend struct {
	Id           int    `json:"id" gorm:"primarykey"`
	UserId       int    `json:"userId" gorm:"column:user_id;index;comment:'用户ID'"`
	FriendUserId int 	`json:"friendUserId" gorm:"index;comment:'好友ID'"`
	Remark       string `json:"remark" gorm:"type:varchar(50);comment:'好友备注"`
	FriendUser  SysUser `json:"friendUser" gorm:"foreignKey:friend_user_id;references:user_id"`
	models.ModelTime
}

func (ImUserFriend) TableName() string {
	return "im_user_friend"
}

type ImUserFriendRequest struct {
	Id           int    	`json:"id" gorm:"primarykey"`
	FromUserId    int    	`json:"fromUserId" gorm:"column:from_user_id;comment:'用户id'"`
	ToUserId      int    	`json:"toUserId" gorm:"column:to_user_id;comment:'要添加的用户id'"`
	HandleResult  int32  	`json:"handleResult" gorm:"column:handle_result;comment: 处理结果：1等待验证 2通过 3拒绝"`
	ReqMsg        string 	`json:"reqMsg" gorm:"column:req_msg;size:255;comment:'好友申请消息'"`
	HandleMsg     string    `json:"handleMsg" gorm:"column:handle_msg;size:255;comment:'好友申请处理消息'"`
	HandleTime    time.Time `json:"handleTime" gorm:"column:handle_time;comment:'好友申请处理时间'"`
	models.ModelTime
}

func (ImUserFriendRequest) TableName() string {
	return "im_friend_request"
}

type ImBlack struct {
	Id           int    	`json:"id" gorm:"primarykey"`
	UserId    		int    `json:"userId" gorm:"comment:'用户Id'"`
	BlockUserId     int    `json:"blockUserId" gorm:"comment:'被拉黑用户Id'"`
	models.ModelTime
}

func (ImBlack) TableName() string {
	return "im_black"
}


