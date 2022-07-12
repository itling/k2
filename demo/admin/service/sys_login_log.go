package service

import (
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"

	"admin/models"
	"admin/service/dto"
)

type SysLoginLog struct {
	service.Service
}

// GetPage 获取SysLoginLog列表
func (e *SysLoginLog) GetPage(c *dto.SysLoginLogGetPageReq, list *[]models.SysLoginLog, count *int64) error {
	var data models.SysLoginLog

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.OrderDest("created_at", true),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysLoginLog对象
func (e *SysLoginLog) Get(d *dto.SysLoginLogGetReq, model *models.SysLoginLog) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Remove 删除SysLoginLog
func (e *SysLoginLog) Remove(c *dto.SysLoginLogDeleteReq) error {
	var data models.SysLoginLog

	db := e.Orm.Delete(&data, c.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}
