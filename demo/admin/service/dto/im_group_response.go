package dto

import "time"

type GroupResponse struct {
	Uuid      string    `json:"uuid"`
	GroupId   int       `json:"groupId"`
	CreatedAt time.Time `json:"createAt"`
	Name      string    `json:"name"`
	Notice    string    `json:"notice"`
}
