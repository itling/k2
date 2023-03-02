package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kingwel-xie/k2/common/api"

	"admin/models"
	"admin/service"
	"admin/service/dto"
)


type ImUserFriend struct {
	api.Api
}


// CreateFriendReq
// @Summary 创建好友申请记录
// @Description 创建好友申请记录
// @Tags  即时通信
// @Param fromUserId query int true "当前用户id"
// @Param toUserId query int true "要添加的用户id"
// @Param reqMsg query string false "添加好友请求消息"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/friend-req [post]
// @Security Bearer
func (e ImUserFriend) CreateFriendReq(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendRequestPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.CreateFriendReq(&req)
	if err != nil {
		e.Error(err)
		return
	}

	e.OK(req.GetId(), "好友申请记录创建成功")
}

// GetFriendRequestPage
// @Summary 获取好友申请列表
// @Description 获取好友申请列表
// @Tags 即时通信
// @Param fromUserId query int false "来源用户ID"
// @Param toUserId query int false "被添加的用户ID"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/friend-req-list [get]
// @Security Bearer
func (e ImUserFriend) GetFriendRequestPage(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendRequestPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]models.ImUserFriendRequest, 0)
	var count int64

	err = s.GetFriendRequestPage(&req, &list, &count)
	if err != nil {
		e.Error(err)
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "好友申请列表查询成功")
}

// GetFriendPage
// @Summary 获取好友列表
// @Description 获取好友列表
// @Tags 即时通信
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/friend-list [get]
// @Security Bearer
func (e ImUserFriend) GetFriendPage(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]models.ImUserFriend, 0)
	var count int64

	err = s.GetFriendPage(&req, &list, &count)
	if err != nil {
		e.Error(err)
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "好友列表查询成功")
}


// GetMessages
// @Summary 获取聊天记录
// @Description 获取聊天记录
// @Tags 即时通信
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/message-list [get]
// @Security Bearer
func (e ImUserFriend) GetMessages(c *gin.Context) {
	s := service.ImMessageService{}
	req := dto.ImMessageRequest{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]dto.ImMessageResponse, 0)

	err = s.GetMessages(&req, &list)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(list,"消息查询成功")
}





// AcceptFriendReq
// @Summary 接受好友请求
// @Description 接受好友请求
// @Tags 即时通信
// @Param id query int true "好友请求记录主键"
// @Param toUserId query int true "用户ID"
// @Param fromUserId query int true "即将被添加的好友用户ID"
// @Param handleMsg query string false "接受消息"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/accept-friend-req [put]
// @Security Bearer
func (e ImUserFriend) AcceptFriendReq(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendRequestPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	err = s.AcceptFriendReq(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "接受好友请求成功")
}



// RejectFriendReq
// @Summary 拒绝好友请求
// @Description 拒绝好友请求
// @Tags 即时通信
// @Param id query int true "好友请求记录主键"
// @Param toUserId query int false "用户ID"
// @Param fromUserId query int false "被拒绝的好友用户ID"
// @Param hanleMsg query string false "拒绝消息"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/reject-friend-req [put]
// @Security Bearer
func (e ImUserFriend) RejectFriendReq(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendRequestPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	err = s.RejectFriendReq(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "拒绝好友请求成功")
}

// UpdateFriendRemark
// @Summary 更新好友备注
// @Description 更新好友备注
// @Tags 即时通信
// @Param id query int true "好友记录主键"
// @Param userId query int false "用户ID"
// @Param friendUserId query int false "好友用户ID"
// @Param remark query string false "备注"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/update-friend-remark [put]
// @Security Bearer
func (e ImUserFriend) UpdateFriendRemark(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	err = s.UpdateFriendRemark(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "更新好友备注成功")
}


// RemoveUserFriend
// @Summary 删除好友
// @Description 删除好友
// @Tags 即时通信
// @Param userId query int false "用户ID"
// @Param friendUserId query int false "好友用户ID"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/delete-friend [delete]
// @Security Bearer
func (e ImUserFriend) RemoveUserFriend(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImUserFriendPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	err = s.RemoveUserFriend(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "删除好友成功")
}



// BlockUserFriend
// @Summary 拉黑好友
// @Description 拉黑好友
// @Tags 即时通信
// @Param userId query int true "用户ID"
// @Param blockUserId query int true "被拉黑用户ID"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/block-friend [post]
// @Security Bearer
func (e ImUserFriend) BlockUserFriend(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImBlackPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	err = s.BlockUserFriend(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "拉黑好友成功")
}



// CancelBlockUserFriend
// @Summary 解除拉黑好友
// @Description 解除拉黑好友
// @Tags 即时通信
// @Param userId query int true "用户ID"
// @Param blockUserId query int true "被拉黑用户ID"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/im/cancel-block-friend [delete]
// @Security Bearer
func (e ImUserFriend) CancelBlockUserFriend(c *gin.Context) {
	s := service.ImUserFriendService{}
	req := dto.ImBlackPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	err = s.CancelBlockUserFriend(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "解除拉黑好友成功")
}


