package dto

import (
	"time"
	"admin/models"
	"admin/constant"
	"github.com/kingwel-xie/k2/common/dto"
)


type ImUserFriendPageReq struct {
	dto.Pagination `search:"-"`
	Id           int    `form:"id"  search:"-" `
	UserId       int    `form:"userId" search:"-"`
	FriendUserId int    `form:"friendUserId" search:"-"`
	Remark       string `form:"remark" search:"-"`
	ImUserFriendOrder
}

type ImUserFriendOrder struct {
}

func (m *ImUserFriendPageReq) GetId() interface{} {
	return m.Id
}

func (m *ImUserFriendPageReq) GetNeedSearch() interface{} {
	return *m
}


func (m *ImUserFriendPageReq) Generate(model *models.ImUserFriend) {
	if  m.Id != 0 {
		model.Id = m.Id
	}
	model.UserId = m.UserId
	model.FriendUserId = m.FriendUserId
	model.Remark = m.Remark
}


type ImUserFriendRequestPageReq struct {
	dto.Pagination `search:"-"`
	Id            int    	`form:"id" search:"-"`
	FromUserId    int    	`form:"fromUserId" search:"-"`
	ToUserId      int    	`form:"toUserId" search:"-"`
	HandleResult  int32  	`form:"handleResult" search:"-"`
	ReqMsg        string 	`form:"reqMsg" search:"-"`
	HandleMsg     string    `form:"handleMsg" search:"-"`
	HandleTime    time.Time `form:"handleTime" search:"-"`
	ImUserFriendRequestOrder
}

type ImUserFriendRequestOrder struct {
}


func (m *ImUserFriendRequestPageReq) GetId() interface{} {
	return m.Id
}


func (m *ImUserFriendRequestPageReq) GetNeedSearch() interface{} {
	return *m
}

func (m *ImUserFriendRequestPageReq) Generate(model *models.ImUserFriendRequest) {
	if  m.Id != 0 {
		model.Id = m.Id
	}
	model.FromUserId = m.FromUserId
	model.ToUserId = m.ToUserId
	model.HandleResult = constant.FRIEND_REQ_WAITING
	model.ReqMsg = m.ReqMsg
	model.HandleMsg = m.HandleMsg
	model.HandleTime = m.HandleTime
}

type ImBlackPageReq struct {
	Id           int     `form:"id" search:"-";`
	UserId    		int  `form:"userId" search:"-"`
	BlockUserId     int  `form:"blockUserId" search:"-"`
	ImBlackPageReqOrder
}


type ImBlackPageReqOrder struct {
}


func (m *ImBlackPageReq) GetId() interface{} {
	return m.Id
}


func (m *ImBlackPageReq) GetNeedSearch() interface{} {
	return *m
}

func (m *ImBlackPageReq) Generate(model *models.ImBlack) {
	if  m.Id != 0 {
		model.Id = m.Id
	}
	model.UserId = m.UserId
	model.BlockUserId = m.BlockUserId
}