package dto

import (
	"admin/models"
)

type SearchResponse struct {
	User  models.SysUser `json:"user"`
	Group models.ImGroup `json:"group"`
}
