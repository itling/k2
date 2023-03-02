package service

import (
	"admin/constant"
	"time"
	"admin/models"
	"admin/service/dto"
	"errors"
	"github.com/kingwel-xie/k2/common/service"
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"gorm.io/gorm"
)

type ImUserFriendService struct {
	service.Service
}


// 创建好友申请记录
func (e *ImUserFriendService) CreateFriendReq(c *dto.ImUserFriendRequestPageReq) error {
	if c.FromUserId!=e.Identity.UserId{
		return k2Error.ErrPermissionDenied
	}

	var friend *models.SysUser
	e.Orm.First(&friend, "user_id = ?", c.ToUserId)
	if friend.UserId == 0 {
		return errors.New("用户不存在 ")
	}

	//判断是否在黑名单中
	var data_black = models.ImBlack{}
	e.Orm.First(&data_black, "user_id = ? and block_user_id=?",c.ToUserId,c.FromUserId)
	if data_black.Id != 0 {
		return errors.New("对方已将你拉入黑名单")
	}

	//判断是否已有申请记录(过滤状态为等待验证的，以前可能删除再添加过)
	var data_friend = models.ImUserFriendRequest{}
	e.Orm.First(&data_friend, "from_user_id = ? and to_user_id=? and handle_result=1",c.FromUserId,c.ToUserId)
	if data_friend.Id != 0 {
		return errors.New("好友申请记录已存在")
	}

	var data models.ImUserFriendRequest
	c.Generate(&data)
	c.FromUserId= e.Identity.UserId
	return e.Orm.Create(&data).Error
}


// 好友列表
func (e *ImUserFriendService) GetFriendPage(c *dto.ImUserFriendPageReq, list *[]models.ImUserFriend, count *int64) error {
	var err error
	var data models.ImUserFriend

	err = e.Orm.Model(&data).Preload("FriendUser").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).Where("user_id",e.Identity.UserId).Order(constant.OrderByCreatedAtDesc).Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// 好友申请列表
func (e *ImUserFriendService) GetFriendRequestPage(c *dto.ImUserFriendRequestPageReq, list *[]models.ImUserFriendRequest, count *int64) error {
	var data models.ImUserFriendRequest
	tx := e.Orm.Model(&data).Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).Where("to_user_id = ? or from_user_id = ?", e.Identity.UserId,e.Identity.UserId)
	return tx.Order(constant.OrderByCreatedAtDesc).Find(list).Limit(-1).Offset(-1).
	Count(count).Error
}

// 通过好友申请
func (e *ImUserFriendService) AcceptFriendReq(c *dto.ImUserFriendRequestPageReq) error {
	if c.ToUserId!=e.Identity.UserId{
		return k2Error.ErrPermissionDenied
	}
	var model = models.ImUserFriendRequest{}
	err := e.Orm.First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	model.HandleMsg=c.HandleMsg;
	model.HandleResult=constant.FRIEND_REQ_PASS;
	model.HandleTime=time.Now();
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		//更新好友申请记录状态
		db := e.Orm.Save(model)
		if db.Error != nil {
			return k2Error.ErrDatabase.Wrap(db.Error)
		}
		if db.RowsAffected == 0 {
			return k2Error.ErrPermissionDenied
		}
		//为自己添加好友记录
		var friend = &dto.ImUserFriendPageReq{
			UserId:model.ToUserId,
			FriendUserId:model.FromUserId,
		}
		err:=e.addFriend(friend)
		if err!=nil {
			return err
		}
		//为对方添加好友记录
		var friendReverse = &dto.ImUserFriendPageReq{
			UserId:model.FromUserId,
			FriendUserId:model.ToUserId,
		}
		return e.addFriend(friendReverse)
	})
	
}


//添加好友
func (e *ImUserFriendService) addFriend(c *dto.ImUserFriendPageReq) error {
	
	var data = new(models.ImUserFriend)
	c.Generate(data)
	var userFriendQuery *models.ImUserFriend
	e.Orm.First(&userFriendQuery, "user_id = ? and friend_user_id = ?", c.UserId, c.FriendUserId)
	//重复记录不抛出错误
	if userFriendQuery.UserId != 0 {
		return nil
	}
	err := e.Orm.Create(data).Error
	return err
}

//设置好友备注
func (e *ImUserFriendService) UpdateFriendRemark(c *dto.ImUserFriendPageReq) error {
	if c.UserId!=e.Identity.UserId{
		return k2Error.ErrPermissionDenied
	}
	var model = models.ImUserFriend{}
	err := e.Orm.First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	model.Remark=c.Remark;
	db := e.Orm.Save(model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// 拒绝好友申请
func (e *ImUserFriendService) RejectFriendReq(c *dto.ImUserFriendRequestPageReq) error {
	if c.ToUserId!=e.Identity.UserId{
		return k2Error.ErrPermissionDenied
	}
	var model = models.ImUserFriendRequest{}
	err := e.Orm.First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	model.HandleMsg=c.HandleMsg;
	model.HandleResult=constant.FRIEND_REQ_REJECT;
	model.HandleTime=time.Now();
	db := e.Orm.Save(model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

func(e *ImUserFriendService) checkDataPerm(userId,friendId int) error{
	if userId!=e.Identity.UserId{
		return k2Error.ErrPermissionDenied
	}
	var friend *models.SysUser
	e.Orm.First(&friend, "user_id = ?", friendId)
	if friend.UserId == 0 {
		return k2Error.ErrCodeNotFound
	}
	return nil
}

//删除好友
func (e *ImUserFriendService)  RemoveUserFriend(d *dto.ImUserFriendPageReq) error {
	err := e.checkDataPerm(d.UserId,d.FriendUserId)
	if err!=nil{
		return err
	}

	var data = models.ImUserFriend{}
	e.Orm.First(&data, "user_id = ? and friend_user_id=?", e.Identity.UserId,d.FriendUserId)
	if data.Id == 0 {
		return k2Error.ErrCodeNotFound
	}

	db := e.Orm.Delete(&data)
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}


//拉黑好友
func (e *ImUserFriendService)  BlockUserFriend(d *dto.ImBlackPageReq) error {
	//判断被拉黑用户是否在用户好友列表，在的话先删除好友
	err := e.checkDataPerm(d.UserId,d.BlockUserId)
	if err!=nil{
		return err
	}

	var data_friend = models.ImUserFriend{}
	e.Orm.First(&data_friend, "user_id = ? and friend_user_id=?", e.Identity.UserId,d.BlockUserId)
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		//删除好友
		if data_friend.Id != 0 {
			db := e.Orm.Delete(data_friend)
			if db.Error != nil {
				return k2Error.ErrDatabase.Wrap(db.Error)
			}
			if db.RowsAffected == 0 {
				return k2Error.ErrPermissionDenied
			}
		}
		//添加黑名单
		var data = new(models.ImBlack)
		d.Generate(data)
		var userBlack *models.ImBlack
		e.Orm.First(&userBlack, "user_id = ? and block_user_id = ?", e.Identity.UserId, d.BlockUserId)
		if userBlack.UserId != 0 {
			return errors.New("该用户被拉黑")
		}
		return e.Orm.Create(data).Error
	})
}

//解除拉黑用户
func (e *ImUserFriendService)  CancelBlockUserFriend(d *dto.ImBlackPageReq) error {
	err := e.checkDataPerm(d.UserId,d.BlockUserId)
	if err!=nil{
		return err
	}
	var data = models.ImBlack{}
	err = e.Orm.First(&data, "user_id = ? and block_user_id=?",d.UserId,d.BlockUserId).Error
	if err != nil {
		return k2Error.ErrCodeNotFound.Wrap(err)
	}

	if data.UserId!=e.Identity.UserId{
		return k2Error.ErrPermissionDenied
	}
	db := e.Orm.Delete(&data)
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

