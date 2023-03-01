package dto

import "time"

type MessageResponse struct {
	Id           int       `json:"id" gorm:"primarykey"`
	FromUserId   int       `json:"fromUserId" gorm:"index"`
	ToUserId     int       `json:"toUserId" gorm:"index"`
	Content      string    `json:"content" gorm:"type:varchar(2500)"`
	ContentType  int16     `json:"contentType" gorm:"comment:'消息内容类型：1文字，2语音，3视频'"`
	CreatedAt    time.Time `json:"createAt"`
	FromUsername string    `json:"fromUsername"`
	ToUsername   string    `json:"toUsername"`
	Avatar       string    `json:"avatar"`
	Url          string    `json:"url"`
}
