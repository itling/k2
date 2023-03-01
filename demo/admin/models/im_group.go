package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type ImGroup struct {
	Id     int    `json:"id" gorm:"primarykey"`
	UserId int    `json:"userId" gorm:"index;comment:'群主ID'"`
	Name   string `json:"name" gorm:"type:varchar(150);comment:'群名称"`
	Notice string `json:"notice" gorm:"type:varchar(350);comment:'群公告"`
	models.ModelTime
}

func (ImGroup) TableName() string {
	return "im_group"
}
